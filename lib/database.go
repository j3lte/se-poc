package lib

import (
  "log"
  "gopkg.in/mgo.v2"
)

var databaseHandle *mgo.Database = nil

func GetDatabaseHandle()(*mgo.Database) {

  if databaseHandle == nil {
    log.Print("Initing database connection")

    session, err := mgo.Dial("mongodb://localhost")
    if err != nil {
      log.Fatal("Could not open database connection!")
      panic(err)
    }
    defer session.Close()

    databaseHandle = session.DB("sepoc")

  }

  return databaseHandle
}