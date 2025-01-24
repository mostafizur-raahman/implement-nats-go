package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	defer nc.Close()

	// Subject to listen for requests
	subject := "example.request"

	// Subscribe to the subject
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received request: %s\n", string(msg.Data))

		// Process the request (e.g., echo the message)
		response := "Processed: " + string(msg.Data)

		// Send the response
		err := msg.Respond([]byte(response))
		if err != nil {
			log.Fatal("Failed to send response:", err)
		}
	})
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}

	log.Printf("Listening for requests on subject: %s\n", subject)

	// Keep the program running
	select {}
}
