package entity

import (
	"time"
)

type Person struct {
	ID        uint   `gorm:"primaryKey;not null;autoIncrement;" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       string `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
