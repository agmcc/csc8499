package metrics

import (
	"context"
	"math"
	"time"

	"k8s.io/klog/v2"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const (
	NetworkLatencyQuery = "network_latency"
	ResponseTimeQuery   = "response_time"
	Missing             = math.MaxInt64
)

func NewAPI(address string) v1.API {
	client, err := api.NewClient(api.Config{Address: address})
	if err != nil {
		klog.Fatalf("Error creating client: %v", err)
	}
	return v1.NewAPI(client)
}

func GetNetworkLatencies(ctx context.Context, api v1.API, nodes []string) map[string]int64 {
	klog.Info("Retrieving network latency metrics for nodes: ", nodes)
	vector := executeQuery(ctx, api, NetworkLatencyQuery)
	m := extractMetricsForNodes(vector, nodes)
	return m
}

func GetResponseTimes(ctx context.Context, api v1.API, nodes []string) map[string]int64 {
	klog.Info("Retrieving response time metrics for nodes: ", nodes)
	vector := executeQuery(ctx, api, ResponseTimeQuery)
	m := extractMetricsForNodes(vector, nodes)
	return m
}

func executeQuery(ctx context.Context, api v1.API, query string) model.Vector {
	klog.Info("Executing query: ", query)
	result, warnings, err := api.Query(ctx, query, time.Now())
	if err != nil {
		klog.Warningf("Error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		klog.Warningf("Warnings: %v", warnings)
	}
	vector := result.(model.Vector)
	klog.Info("Result vector: ", vector)
	return vector
}

func extractMetricsForNodes(vector model.Vector, nodes []string) map[string]int64 {
	m := make(map[string]int64)
	for _, n := range nodes {
		m[n] = Missing
	}

	for _, sample := range vector {
		nodeName := string(sample.Metric["host"])
		_, exists := m[nodeName]
		if exists {
			value := float64(sample.Value)
			if !math.IsNaN(value) {
				metric := int64(sample.Value)
				m[nodeName] = metric
				klog.Infof("Extracted value %d for node '%s'", metric, nodeName)
			} else {
				klog.Infof("NaN value for node '%s'", nodeName)
			}
		}
	}
	return m
}
