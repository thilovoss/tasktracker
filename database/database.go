package database

import (
	"fmt"
	"log"
	"tasktracker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(url string) {
	Instance, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		panic("Could not connect to database")
	}
	fmt.Println("Connected to database...")
}

func Migrate() {
	Instance.AutoMigrate(&models.Task{}, models.Category{})
	fmt.Println("Database Migration completed...")
}
