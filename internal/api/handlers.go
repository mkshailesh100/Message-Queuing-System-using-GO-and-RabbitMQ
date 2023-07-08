package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkshailesh100/message-queue-system/internal/db"
	"github.com/mkshailesh100/message-queue-system/internal/messaging"
	"github.com/mkshailesh100/message-queue-system/pkg/models"
)
type Product struct {
	ID					uint 	`json:"id"`
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      string `json:"product_images"`
	ProductPrice       float64  `json:"product_price"`
}
func CreateProduct(c *gin.Context) {
    // Parse the request body and bind it to the Product struct
    var product Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	log.Println("Product details part1:", product)
    // Connect to the database
    db, err := db.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
        return
    }

    // Create a new product record in the database
    newProduct := models.Product{
		ID:					product.ID,
        UserID:             product.UserID,
        ProductName:        product.ProductName,
        ProductDescription: product.ProductDescription,
        ProductImages:      product.ProductImages, 
        ProductPrice:       product.ProductPrice,
    }
    if err := db.Create(&newProduct).Error; err != nil {
        log.Println("Failed to create product:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    // Publish the product ID to the message queue
    log.Println("Product details:", newProduct)
    err = messaging.PublishProductID(int(newProduct.ID))
    if err != nil {
        log.Println("Failed to publish product ID:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish product ID"})
        return
    }

    // Return a success response
    c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})

	go messaging.ConsumeProductID()
	
}
