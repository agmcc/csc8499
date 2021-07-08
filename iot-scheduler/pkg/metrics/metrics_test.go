package metrics

import (
	"context"
	"testing"
)

func TestGetLatencies(t *testing.T) {
	t.Log("Running test")
	nodes := []string{"docker-desktop"}
	latencies := GetLatencies(context.TODO(), "http://localhost:32000", nodes)
	for n, l := range latencies {
		t.Logf("%s: %d", n, l)
	}
}
