package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subject to subscribe to
	subject := "example.subject"

	// Subscribe to the subject
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Subscribed to subject: %s\n", subject)

	// Keep the connection alive to receive messages
	select {}
}
