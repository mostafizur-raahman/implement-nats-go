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

	// Subject and queue group to subscribe to
	subject := "example.queue"
	queueGroup := "my_queue_group"

	// Subscribe to the subject as part of the queue group
	_, err = nc.QueueSubscribe(subject, queueGroup, func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}

	log.Printf("Subscribed to subject '%s' in queue group '%s'\n", subject, queueGroup)

	// Keep the program running
	select {}
}
