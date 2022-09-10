package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
)

// create person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to create person")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	// database.Connector.Omit("id")
	database.Connector.Omit("ID").Create(&person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}
