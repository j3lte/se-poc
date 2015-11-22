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
  BirthDate         int           `json:"birthdate" bson:"birthdate"`
  Addresses         []Address     `json:"addresses" bson:"addresses"`
  Accounts          []Account     `json:"accounts" bson:"accounts"`
  PhoneNumbers      []PhoneNumber `json:"phonenumbers" bson:"phonenumbers"`
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

  // Does the object exist?
  if !bson.IsObjectIdHex(string(p.Id)) {
    p.Id = bson.NewObjectId()
    err := c.Insert(p)

    if err != nil {
      log.Panic("Error whilst inserting person: %v", err)
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