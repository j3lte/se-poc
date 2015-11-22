package people

import (
	"log"
	"encoding/json"
	"net/http"
)

func init() {

  log.Printf("Registering the people routes")

	http.HandleFunc("/api/people", RouteGet)
	http.HandleFunc("/api/people/add", RouteAdd)

}

func RouteAdd(res http.ResponseWriter, req *http.Request) {

  p := GetNewPerson()

  // Decode the JSON against the Person struct
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&p)
  if err != nil {
    log.Panic("Could not properly decode request \n body: %v", req.Body)
  }

  // Save the actual data
  SavePerson(p)

  log.Printf("Adding in new person: %s", p.FirstName + " " + p.LastName)

  res.Header().Set("Content-Type", "application/json")
  res.Write([]byte("\"success\""))

}

func RouteGet(res http.ResponseWriter, req *http.Request) {
  var personID  string    = req.URL.Query().Get("person")
  var data      []byte    = nil
  var err       error     = nil


  if personID != "" {
    result := GetPerson(personID)
    data, err = json.Marshal(result)
  } else {
    results := AllPeople()
    data, err = json.Marshal(results)
  }

  if err != nil {
    log.Panic("Could not serialize data for GET /people")
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(data)
}

func RouteGetOne(_ http.ResponseWriter, req *http.Request) {
  log.Printf("Got request: %s", req.URL.Query().Get("person"))
}