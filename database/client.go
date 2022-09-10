package database

import (
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

// Connector variable used for CRUD operation's
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	// log.Println("Connection was successful!!")
	return nil
}

// migrate database table
func Migrate(table *entity.Person) {
	Connector.AutoMigrate(&table)
	// log.Println("table created successfully")
}
