package main

import (
	"context"
	"log"
	"time"

	"github.com/plgd-dev/go-coap/v2/udp"
)

func main() {
	co, err := udp.Dial(":5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := co.Get(ctx, "/greeting")
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	bytes, err := resp.ReadBody()
	if err != nil {
		log.Printf("Failed to read response: %v", err)
	}
	log.Printf("Response: %s", bytes)
}
