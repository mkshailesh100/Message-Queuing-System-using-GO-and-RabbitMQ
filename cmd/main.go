package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mkshailesh100/message-queue-system/internal/db"
    "github.com/mkshailesh100/message-queue-system/internal/api"
)

func main() {
	// Connect to the database
	dbInstance, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Run database migrations
	err = db.Migrate(dbInstance)
	if err != nil {
		panic("Failed to run database migrations")
	}

	// Create a new Gin router
	r := gin.Default()

	// Define the API endpoint for creating a product
	r.POST("/products", api.CreateProduct)

	// Start the server
	log.Println("Server started at localhost:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to start server: %w", err))
	}
	// Start your application logic here
}
