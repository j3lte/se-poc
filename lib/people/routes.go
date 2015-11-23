package people

import (
	"log"
	"encoding/json"
	"net/http"
  "github.com/BetaBugish/se-poc/lib"
)

func init() {

  log.Printf("Registering the people routes")

	http.HandleFunc("/api/people", RouteGet)
  http.HandleFunc("/api/people/add", RouteAdd)
  http.HandleFunc("/api/people/save", RouteSave)
  http.HandleFunc("/api/people/new", RouteNew)
	http.HandleFunc("/api/people/map", RouteMap)
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

// Handles the save request emitted when creating or updating a person
func RouteSave(res http.ResponseWriter, req *http.Request) {

  p := GetNewPerson()

  // Decode the JSON against the Person struct
  decoder := json.NewDecoder(req.Body)
  err := decoder.Decode(&p)
  if err != nil {
    log.Panic("Could not properly decode request \n body: %v", req.Body)
  }

  // Save the actual data
  SavePerson(p)

  log.Printf("Updating person: %s", p.FirstName + " " + p.LastName)

  res.Header().Set("Content-Type", "application/json")
  res.Write([]byte("\"success\""))
}

// Outputs a blank person, used to set up the structure in Angular
func RouteNew(res http.ResponseWriter, req *http.Request) {
  p := GetNewPerson()
  data, err := json.Marshal(p)

  if err != nil {
    log.Panic("Could not output a blank canvas for the person object")
    panic(err)
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(data)
}

// Generates a map for use with vis
func RouteMap(res http.ResponseWriter, req *http.Request) {
  var personID  string    = req.URL.Query().Get("person")
  var counter int = 0
  var networkMap *lib.NetworkMap

  result := GetPerson(personID)
  counter, networkMap = result.ToNetworkMap(counter, make(map[string]int))

  data, err := json.Marshal(networkMap)

  if err != nil {
    log.Panic("Could not output a blank canvas for the person object")
    panic(err)
  }

  res.Header().Set("Content-Type", "application/json")
  res.Write(data)
}

