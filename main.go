package main

import (
	"net/http"
	"log"
	"github.com/BetaBugish/se-poc/lib"
	_ "github.com/BetaBugish/se-poc/lib/people"
)

var port string = ":1234"

func main() {

	lib.GetDatabaseHandle()

	log.Print("Registering webserver routes");

	/* Routing */
	http.Handle("/", http.FileServer(http.Dir("public/")))

	log.Printf("Starting webserver on port %s", port);
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
}