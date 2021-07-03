package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-ping/ping"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

var httpDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name:       "network_latency_seconds",
	Help:       "RTT latency to edge nodes.",
	Objectives: map[float64]float64{0.5: 0.05, 0.99: 0.001},
	// TODO rename "host" to "node" (and also in go-server)
}, []string{"host", "status"})

var hosts []string
var interval time.Duration
var port int
var node string
var pingCount int
var timeout time.Duration

func main() {
	initEnv()
	log.Printf("Checking latency for hosts %s every %s\n", hosts, interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	done := make(chan bool)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	go func() {
		r := mux.NewRouter()
		r.Handle("/metrics", promhttp.Handler())
		log.Printf("Starting server on port %d...\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
	}()

	checkLatency()
	for {
		select {
		case <-ticker.C:
			checkLatency()
		case <-done:
			log.Println("Exiting")
			return
		}
	}

}

func checkLatency() {
	log.Printf("Checking hosts: %s\n", hosts)
	for _, h := range hosts {
		log.Printf("Pinging %s\n", h)
		stats := pingHost(h)
		if stats != nil && len(stats.Rtts) > 0 {
			log.Printf("RTTs %s\n", stats.Rtts)
			for _, rtt := range stats.Rtts {
				httpDuration.WithLabelValues(node, "success").Observe(rtt.Seconds())
			}
		} else {
			log.Printf("Failed to ping %s\n", h)
			httpDuration.WithLabelValues(node, "fail").Observe(0)
		}
	}
}

func pingHost(ip string) *ping.Statistics {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Println(err)
		return nil
	}
	pinger.Count = pingCount
	pinger.Timeout = timeout
	pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		log.Println(err)
		return nil
	}
	return pinger.Statistics()
}

func initEnv() {
	viper.SetEnvPrefix("agent")
	viper.AutomaticEnv()

	viper.SetDefault("ping_count", 3)
	viper.SetDefault("interval", time.Duration(15)*time.Second)
	viper.SetDefault("timeout", time.Duration(5)*time.Second)
	viper.SetDefault("port", 9000)
	viper.SetDefault("node", "Unknown")

	pingCount = viper.GetInt("ping_count")
	interval = viper.GetDuration("interval")
	timeout = viper.GetDuration("timeout")
	hosts = viper.GetStringSlice("hosts")
	port = viper.GetInt("port")
	node = viper.GetString("node")
}
