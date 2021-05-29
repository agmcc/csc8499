package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
	fmt.Println("Endpoint Hit: greeting")
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
