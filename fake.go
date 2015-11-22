package main

/* This file is meant for generating test data */

import (
  "github.com/BetaBugish/se-poc/lib/people"
  "github.com/icrowley/fake"
  "log"
)

func main() {
  log.Print("Generating fake data")

  for i := 0; i < 10; i++ {
    p := people.GetNewPerson()

    p.FirstName = fake.FirstName()
    p.LastName = fake.LastName()
    p.Email = fake.EmailAddress()

    // Fake accounts
    for j := 0; j < 3; j++ {
      a := people.GetNewAccount()

      a.Password = fake.SimplePassword()
      a.UserName = fake.UserName()

      p.Accounts = append(p.Accounts, *a)
    }

    // Fake addresses
    for j := 0; j < 3; j++ {
      a := people.GetNewAddress()

      a.Streetname = fake.Street()
      a.City = fake.City()
      a.State = fake.State()
      a.Housenumber = fake.DigitsN(2)
      a.Postal = fake.Zip()

      p.Addresses = append(p.Addresses, *a)
    }

    // Fake phone numbers
    for j := 0; j < 3; j++ {
      p.PhoneNumbers = append(p.PhoneNumbers, fake.Phone(), fake.Phone(), fake.Phone())
    }

    people.SavePerson(p)
  }

  log.Print("Done generating fake data")
}