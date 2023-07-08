package messaging

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)
func PublishProductID(productID int) error {
    // Connect to RabbitMQ server
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return err
    }
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    if err != nil {
        return err
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
        return err
    }

    // Publish the product ID to the queue
    err = ch.Publish(
        "",        // Exchange
        q.Name,    // Routing key
        false,     // Mandatory
        false,     // Immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(strconv.Itoa(productID)),
        },
    )
    if err != nil {
        return err
    }

    log.Println("Product ID published to the message queue")
    return nil
}
