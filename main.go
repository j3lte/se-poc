package main

import (
	"net/http"
	"log"
	"github.com/BetaBugish/se-poc/lib"
	"github.com/BetaBugish/se-poc/lib/people"
)

func main() {

	lib.GetDatabaseHandle()

	log.Print("Registering webserver routes");

	/* Routing */
	http.Handle("/", http.FileServer(http.Dir("public/")))
	
	log.Print("Starting webserver on port :666");
	err := http.ListenAndServe(":666", nil)
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
}