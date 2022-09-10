package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
)

// create person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to create person")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(&person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

// get person by id
func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to get person by ID")
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
