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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var difficulty = getDifficulty()
var host = hostname()

var httpDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name:       "http_response_time_seconds",
	Help:       "Duration of HTTP requests.",
	Objectives: map[float64]float64{0.5: 0.05, 0.99: 0.001},
}, []string{"path", "host"})

type Load struct {
	Host       string        `json:"host"`
	Difficulty int           `json:"difficulty"`
	Match      string        `json:"match"`
	Elapsed    time.Duration `json:"elapsed"`
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	match := doWork(difficulty)
	elapsed := time.Since(start)
	response := Load{host, difficulty, match, elapsed}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getDifficulty() int {
	difficultyStr, exists := os.LookupEnv("DIFFICULTY")
	if !exists {
		fmt.Println("No difficulty set, defaulting to 0")
		difficultyStr = "0"
	}

	difficulty, err := strconv.Atoi(difficultyStr)
	if err != nil {
		fmt.Println(err)
	}

	if difficulty >= 0 {
		return difficulty
	} else {
		return 0
	}
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

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, host))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(metricsMiddleware)

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/load", loadHandler)

	fmt.Printf("Hostname: %s\n", host)
	fmt.Printf("Difficulty: %d\n", difficulty)
	fmt.Println("Listening on port 8080 ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
