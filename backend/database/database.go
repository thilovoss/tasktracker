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
	err := Instance.AutoMigrate(models.Category{}, &models.Task{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Database Migration completed...")
}
