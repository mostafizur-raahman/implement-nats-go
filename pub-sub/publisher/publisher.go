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
		log.Fatal(err)
	}
	defer nc.Close()

	// Subject to publish to
	subject := "example.subject"

	// Message to send
	message := "Hello, NATS!"

	// Publish the message
	err = nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message: %s\n", message)

	// Wait a bit to ensure the message is sent
	time.Sleep(1 * time.Second)
}
