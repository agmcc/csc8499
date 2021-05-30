package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hello from "+name)
	fmt.Println("Endpoint Hit: greeting")
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
