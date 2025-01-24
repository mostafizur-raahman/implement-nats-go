package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer nc.Close()

	// Subject to publish to
	subject := "example.queue"

	// Publish messages every second
	for i := 1; i <= 5; i++ {
		message := "Message " + string(i)
		err := nc.Publish(subject, []byte(message))
		if err != nil {
			log.Fatal("Failed to publish message:", err)
		}
		log.Printf("Published message: %s\n", message)
		time.Sleep(1 * time.Second)
	}
}
