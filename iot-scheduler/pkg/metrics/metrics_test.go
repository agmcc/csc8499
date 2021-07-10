package metrics

import (
	"context"
	"testing"
)

const (
	Address = "http://localhost:32000"
)

func TestGetNetworkLatencies(t *testing.T) {
	nodes := []string{"docker-desktop", "unknown"}
	api := NewAPI(Address)
	latencies := GetNetworkLatencies(context.TODO(), api, nodes)
	t.Log("Latencies: ", latencies)

}

func TestGetResponseTimes(t *testing.T) {
	nodes := []string{"docker-desktop", "unknown"}
	api := NewAPI(Address)
	responseTimes := GetResponseTimes(context.TODO(), api, nodes)
	t.Log("Response times: ", responseTimes)
}
