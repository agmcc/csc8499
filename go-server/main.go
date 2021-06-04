package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	coap "github.com/plgd-dev/go-coap/v2"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/mux"
)

func loggingMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		log.Printf(r.String())
		next.ServeCOAP(w, r)
	})
}

func greeting(w mux.ResponseWriter, r *mux.Message) {
	name := "<Unknown>"
	host, found := os.LookupEnv("HOST")
	if found {
		name = host
	}
	log.Print("Sending greeting")
	err := w.SetResponse(codes.Content, message.TextPlain, bytes.NewReader([]byte("Hello from "+name)))
	if err != nil {
		log.Printf("Failed to set response: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Handle("/greeting", mux.HandlerFunc(greeting))
	fmt.Println("Listening on UDP port 5688...")
	log.Fatal(coap.ListenAndServe("udp", ":5688", r))
}
