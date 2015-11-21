package people

import (
	"log"
	"net/http"
)

func init() {

	http.Handle("/api/people/add", RouteAdd)
}

func RouteAdd(_ http.ResponseWrite, req *http.Request) {

	req.Write("<h1>Hello!</h1>")
}