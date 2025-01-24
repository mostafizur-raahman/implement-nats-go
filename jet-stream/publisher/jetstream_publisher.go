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

	// Create a stream
	streamName := "MY_STREAM"
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{"example.durable"},
		Storage:  nats.FileStorage,
	})
	if err != nil {
		log.Fatal("Failed to create stream:", err)
	}
	log.Println("Stream created:", streamName)

	// Publish messages to the stream
	for i := 1; i <= 5; i++ {
		message := "Message " + string(i)
		_, err := js.Publish("example.durable", []byte(message))
		if err != nil {
			log.Fatal("Failed to publish message:", err)
		}
		log.Printf("Published message1: %s\n", message)
		time.Sleep(1 * time.Second)
	}
}
