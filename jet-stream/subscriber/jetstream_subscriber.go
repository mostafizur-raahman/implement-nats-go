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

	// Create a JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("Failed to create JetStream context:", err)
	}

	// Create a durable consumer
	streamName := "MY_STREAM"
	consumerName := "MY_CONSUMER"
	_, err = js.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:       consumerName,
		AckPolicy:     nats.AckExplicitPolicy,
		DeliverPolicy: nats.DeliverAllPolicy,
	})
	if err != nil {
		log.Fatal("Failed to create consumer:", err)
	}
	log.Println("Consumer created:", consumerName)

	// Subscribe to the stream using PullSubscribe
	sub, err := js.PullSubscribe("example.durable", consumerName)
	if err != nil {
		log.Fatal("Failed to subscribe:", err)
	}
	defer sub.Unsubscribe()

	log.Println("Listening for messages...")

	// Process messages
	for {
		// Fetch messages (batch size of 1)
		msgs, err := sub.Fetch(1, nats.MaxWait(10*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				log.Println("No messages received. Waiting...")
				continue
			}
			log.Fatal("Failed to fetch messages:", err)
		}

		// Process each message
		for _, msg := range msgs {
			log.Printf("Received message: %s\n", string(msg.Data))

			// Acknowledge the message
			err = msg.Ack()
			if err != nil {
				log.Fatal("Failed to acknowledge message:", err)
			}
		}
	}
}
