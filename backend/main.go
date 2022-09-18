package main

import (
	"tasktracker/database"
	"tasktracker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := "postgres://root:root@localhost:5432/tasktracker"
	database.Connect(dbURL)
	database.Migrate()

	r := gin.Default()

	handler.NewHandler(&handler.Config{
		R: r,
	})

	r.Run()
}
