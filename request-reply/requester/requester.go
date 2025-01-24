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

	// Subject to send the request to
	subject := "example.request"

	// Request message
	request := "Hello, can you process this?"

	// Send the request and wait for a response
	response, err := nc.Request(subject, []byte(request), 2*time.Second)
	if err != nil {
		log.Fatal("Failed to get response:", err)
	}

	log.Printf("Received response: %s\n", string(response.Data))
}
