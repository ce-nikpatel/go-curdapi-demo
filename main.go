package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	//"github.com/ssaumyaranjan7/simple-REST/pkg/controllers"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePerson).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "learning",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}
