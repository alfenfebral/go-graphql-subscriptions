package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	// Simple Publisher
	nc.Publish("payload-subject", []byte("Hello World"))

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
