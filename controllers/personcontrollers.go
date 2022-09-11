package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"rest-go-demo/response"
	"strconv"

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

// update person
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to update person by ID")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

// delete person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to delete person by ID")
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusOK)
	response.Write(w, response.Response{
		Code:    http.StatusOK,
		Message: "person deleted successfully",
	})

}

// get all person list
func GetPersonList(w http.ResponseWriter, r *http.Request) {
	log.Println("start code to get all person list")
	var persons []entity.Person
	database.Connector.Find(&persons).Order("id DESC", true)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}
