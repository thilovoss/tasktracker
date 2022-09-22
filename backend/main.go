package main

import (
	"tasktracker/database"
	"tasktracker/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := "postgres://root:root@localhost:5432/tasktracker"
	database.Connect(dbURL)
	database.Migrate()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	handler.NewHandler(&handler.Config{
		R: r,
	})

	r.Run()
}
