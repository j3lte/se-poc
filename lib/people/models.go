package people

import (
  "github.com/BetaBugish/se-poc/lib"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
)

var collection *mgo.Collection = nil

type Person struct {
  Id           bson.ObjectId `json:"_id" bson:"_id"`
  FirstName    string        `json:"firstname" bson:"firstname"`
  LastName     string        `json:"lastname" bson:"lastname"`
  BirthDate    string        `json:"birthdate" bson:"birthdate"`
  Addresses    []Address     `json:"addresses" bson:"addresses"`
  Accounts     []Account     `json:"accounts" bson:"accounts"`
  PhoneNumbers []PhoneNumber `json:"phonenumbers" bson:"phonenumbers"`
  Relations    []Relation    `json:"relations" bson:"relations"`
}

type Account struct {
  Password string `json:"password" bson:"password"`
  UserName string `json:"username" bson:"username"`
  Link     string `json:"link" bson:"link"`
}

type Address struct {
  Streetname  string `json:"streetname" bson:"streetname"`
  City        string `json:"city" bson:"city"`
  State       string `json:"state" bson:"state"`
  Country     string `json:"country" bson:"country"`
  Housenumber string `json:"housenumber" bson:"housenumber"`
  Postal      string `json:"postal" bson:"postal"`
}

type PhoneNumber struct {
  Number string `json:"number" bson:"number"`
  Type   string `json:"type" bson:"type"`
}

type Relation struct {
  Description string        `json:"description" bson:"description"`
  SubjectType string        `json:"subjecttype" bson:"subjecttype"`
  Subject     bson.ObjectId `json:"subject" bson:"subject"`
}

// Turns the person into a node item for vis
func (p *Person) ToNode(id int) *lib.Node {

  n := lib.GetNewNode()

  n.Label = p.FirstName + " " + p.LastName
  n.Id = id
  n.Group = "person"

  return n
}

// Turns the person into a network map
// To Do: Overhaul this to be more readable
func (p *Person) ToNetworkMap(counter int, checklist map[string]int) (int, *lib.NetworkMap) {
  networkMap := lib.GetNewNetworkMap()
  var exists bool
  var currentPersonID int
  var currentChecklist = checklist

  // Check if the person already exists as a node
  currentPersonID, exists = currentChecklist[p.Id.Hex()]

  if !exists {

    // Add the person as a node
    networkMap.AddNode(p.ToNode(counter))
    currentChecklist[p.Id.Hex()] = counter
    currentPersonID = counter
    counter++
  }

  // Add the person's addresses as nodes
  for _, address := range p.Addresses {
    networkMap.AddNode(address.ToNode(counter))
    networkMap.CreateEdge(currentPersonID, counter)
    counter++
  }

  // Add the person's accounts as nodes
  for _, account := range p.Accounts {
    networkMap.AddNode(account.ToNode(counter))
    networkMap.CreateEdge(currentPersonID, counter)
    counter++
  }

  // Add the person's relations as network maps
  for _, relation := range p.Relations {
    var relationNetworkMap *lib.NetworkMap
    // Holds the vis reference for the matched person as a node
    checklistMatch, exists := checklist[relation.Subject.Hex()]
    p := GetPerson(relation.Subject.Hex())

    // Check if the relation has a node already so we don't re-add it
    if !exists {

      // And add the edge linking the new person to the relation owner
      networkMap.CreateEdge(currentPersonID, counter)

      // Create a new network map from the fetched person
      counter, relationNetworkMap = p.ToNetworkMap(counter, checklist)

      // Make the main network map eat the fetched person's network map
      networkMap.Absorb(relationNetworkMap)

      // And add the edge linking the new person to the relation owner
      networkMap.CreateEdge(currentPersonID, counter)
    } else {

      // No new node was added, but there was a relation
      networkMap.CreateEdge(currentPersonID, checklistMatch)
    }

    counter++
  }

  return counter, networkMap
}

// Turns the account into a node item for vis
func (a *Account) ToNode(id int) *lib.Node {

  n := lib.GetNewNode()

  n.Label = a.UserName
  n.Id = id
  n.Group = "account"

  return n
}

// Turns the account into a node item for vis
func (a *Address) ToNode(id int) *lib.Node {

  n := lib.GetNewNode()

  n.Label = a.Streetname + a.Housenumber
  n.Id = id
  n.Group = "address"

  return n
}

func GetNewPerson() *Person {
  return &Person{}
}

func GetNewAccount() *Account {
  return &Account{}
}

func GetNewAddress() *Address {
  return &Address{}
}

func GetNewPhoneNumber() *PhoneNumber {
  return &PhoneNumber{}
}

func getCollection() *mgo.Collection {
  if collection == nil {
    collection = lib.GetDatabaseHandle().C("people")
  }

  return collection
}

func SavePerson(p *Person) {
  c := getCollection()

  log.Printf("Saving person: %s", p.Id.Hex())

  log.Printf("Person has valid objectid: %s", bson.IsObjectIdHex(p.Id.Hex()))

  // Does the object exist?
  if !bson.IsObjectIdHex(p.Id.Hex()) {
    p.Id = bson.NewObjectId()
    err := c.Insert(p)

    log.Printf("Inserted person: %v", p.Id)
    if err != nil {
      log.Panic("Error whilst inserting person: %v", err)
    }
  } else {
    err := c.UpdateId(p.Id, p)

    log.Printf("Saved person: %v", p.Id)
    if err != nil {
      log.Panic("Error whilst updating person: %v", err)
    }
  }
}

func AllPeople() []Person {
  c := getCollection()
  var results []Person

  c.Find(nil).All(&results)

  return results
}

func GetPerson(id string) *Person {
  c := getCollection()
  result := GetNewPerson()

  c.FindId(bson.ObjectIdHex(id)).One(&result)

  return result
}
