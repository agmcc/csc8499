package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Load struct {
	Host       string `json:"host"`
	Difficulty int    `json:"difficulty"`
	Match      string `json:"match"`
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	difficulty := getDifficulty()
	match := doWork(difficulty)
	host := hostname()
	response := Load{host, difficulty, match}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getDifficulty() int {
	difficultyStr, exists := os.LookupEnv("DIFFICULTY")
	if !exists {
		difficultyStr = "4"
	}

	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		fmt.Println(err)
	}
	return difficulty
}

func doWork(difficulty int) string {
	target := strings.Repeat("0", difficulty)
	rand.Seed(time.Now().UTC().UnixNano())
	start := rand.Int63()

	for true {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, start)
		hash := md5.Sum(buf.Bytes())
		encoded := hex.EncodeToString(hash[:])
		slice := encoded[:difficulty]
		if slice == target {
			return encoded
		}
		start = start + 1
	}
	return ""
}

func hostname() string {
	name := "Unknown"
	host, found := os.LookupEnv("HOST")
	if found {
		name = host
	}
	return name
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/load", loadHandler)
	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
