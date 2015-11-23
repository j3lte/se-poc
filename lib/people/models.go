package people

import (
  "log"
  "github.com/BetaBugish/se-poc/lib"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

var collection *mgo.Collection = nil

type Person struct {
  Id                bson.ObjectId `json:"_id" bson:"_id"`
  FirstName         string        `json:"firstname" bson:"firstname"`
  LastName          string        `json:"lastname" bson:"lastname"`
  BirthDate         string        `json:"birthdate" bson:"birthdate"`
  Addresses         []Address     `json:"addresses" bson:"addresses"`
  Accounts          []Account     `json:"accounts" bson:"accounts"`
  PhoneNumbers      []PhoneNumber `json:"phonenumbers" bson:"phonenumbers"`
  Relations         []Relation    `json:"relations" bson:"relations"`
}

type Account struct {
  Password      string    `json:"password" bson:"password"`
  UserName      string    `json:"username" bson:"username"`
  Link          string    `json:"link" bson:"link"`
}

type Address struct {
  Streetname    string    `json:"streetname" bson:"streetname"`
  City          string    `json:"city" bson:"city"`
  State         string    `json:"state" bson:"state"`
  Country       string    `json:"country" bson:"country"`
  Housenumber   string    `json:"housenumber" bson:"housenumber"`
  Postal        string    `json:"postal" bson:"postal"`
}

type PhoneNumber struct {
  Number        string    `json:"number" bson:"number"`
  Type          string    `json:"type" bson:"type"`
}

type Relation struct {
  Description   string          `json:"description" bson:"description"`
  SubjectType   string          `json:"subjecttype" bson:"subjecttype"`
  Subject       bson.ObjectId   `json:"subject" bson:"subject"`
}

// Turns the person into a node item for vis
func (p *Person)ToNode(id int, group string)(*lib.Node) {

    n := lib.GetNewNode()

    n.Label = p.FirstName + " " + p.LastName
    n.Id = id
    n.Group = group

    return n
}

// Turns the account into a node item for vis
func (a *Account)ToNode(id int)(*lib.Node) {

    n := lib.GetNewNode()

    n.Label = a.UserName
    n.Id = id
    n.Group = "account"

    return n
}

// Turns the account into a node item for vis
func (a *Address)ToNode(id int)(*lib.Node) {

    n := lib.GetNewNode()

    n.Label = a.Streetname + a.Housenumber
    n.Id = id
    n.Group = "address"

    return n
}

func GetNewPerson()(*Person) {
  return &Person{}
}

func GetNewAccount()(*Account) {
  return &Account{}
}

func GetNewAddress()(*Address) {
  return &Address{}
}

func GetNewPhoneNumber()(*PhoneNumber) {
  return &PhoneNumber{}
}

func getCollection() (*mgo.Collection) {
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

func AllPeople()([]Person) {
  c := getCollection()
  var results []Person

  c.Find(nil).All(&results)

  return results
}

func GetPerson(id string)(*Person) {
  c := getCollection()
  result := GetNewPerson()

  c.FindId(bson.ObjectIdHex(id)).One(&result)

  return result
}