package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	name := "<Unknown>"
	host, found := os.LookupEnv("HOST")
	if found {
		name = host
	}
	fmt.Fprintf(w, "Hello from "+name)
	fmt.Println("Endpoint Hit: greeting")
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
