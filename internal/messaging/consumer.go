package messaging

import (
	"log"
	"strconv"
	"strings"

	"github.com/mkshailesh100/message-queue-system/internal/compression"
	"github.com/mkshailesh100/message-queue-system/internal/db"
	"github.com/mkshailesh100/message-queue-system/pkg/models"
	"github.com/streadway/amqp"
)
func ConsumeProductID() {
	log.Println("Inside consumer")
    // Connect to RabbitMQ server
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal(err)
    }
    defer ch.Close()

    // Declare a queue
    q, err := ch.QueueDeclare(
        "product_ids", // Queue name
        false,         // Durable
        false,         // Delete when unused
        false,         // Exclusive
        false,         // No-wait
        nil,           // Arguments
    )
    if err != nil {
        log.Fatal(err)
    }

    // Consume messages from the queue
    msgs, err := ch.Consume(
        q.Name, // Queue name
        "",     // Consumer name
        true,   // Auto-acknowledge
        false,  // Exclusive
        false,  // No-local
        false,  // No-wait
        nil,    // Arguments
    )
    if err != nil {
        log.Fatal(err)
    }

	

    // Process incoming messages
    for msg := range msgs {
		log.Println("consumer msg is ", msg)
        productID, err := strconv.Atoi(string(msg.Body))
        if err != nil {
            log.Println("Failed to convert product ID:", err)
        } else {
            // Process the product ID
            log.Println("Product ID processed:", productID)

            // Get the product details from the database
            dbConn, err := db.ConnectDB()
            if err != nil {
                log.Println("Failed to connect to the database:", err)
                continue
            }

            var product models.Product
            result := dbConn.First(&product, productID)
            if result.Error != nil {
                log.Println("Failed to fetch product from the database:", result.Error)
                continue
            }
			imagesArray := strings.Split(product.ProductImages, ",")

			// Remove leading/trailing whitespaces from each element
			for i := 0; i < len(imagesArray); i++ {
				imagesArray[i] = strings.TrimSpace(imagesArray[i])
			}
			
            compressedPaths, err := compression.DownloadAndCompressImages(imagesArray)
            if err != nil {
                log.Println("Failed to download and compress images:", err)
                continue
            }

            // Update the database with the compressed image paths
			compressedString :=  strings.Join(compressedPaths, ",")
            product.CompressedProductImages = compressedString
            result = dbConn.Save(&product)
            if result.Error != nil {
                log.Println("Failed to update product in the database:", result.Error)
            }
        }
    }
}
