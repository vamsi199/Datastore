package main

import (
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"

	"net/http"
	"fmt"
	"log"
)
func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080",nil))

}
type Employee struct {
	Name string
	Role string

}
func healthCheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"ok")
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	e1 := Employee{
		Name: "vamsi" ,
		Role: "manager" ,
	}
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "employee" ,nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var e2 Employee
	if err = datastore.Get(c, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"put and get value %q" ,e2.Name)

}