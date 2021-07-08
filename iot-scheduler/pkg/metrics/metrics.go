package metrics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/klog/v2"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const (
	QueryTmpl = "ceil(avg without (instance) (network_latency_seconds{status='success', quantile='0.5', host=~'%s'}) * 1000)"
)

func GetLatencies(ctx context.Context, address string, nodes []string) map[string]int64 {
	klog.Infof("Retrieving latency metrics for nodes: %s", nodes)

	// Create client
	client, err := api.NewClient(api.Config{Address: address})
	if err != nil {
		klog.Fatalf("Error creating client: %v", err)
	}
	v1api := v1.NewAPI(client)

	// Query
	query := fmt.Sprintf(QueryTmpl, strings.Join(nodes, "|"))
	klog.Info("Executing query: ", query)
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		klog.Warningf("Error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		klog.Warningf("Warnings: %v", warnings)
	}

	vector := result.(model.Vector)

	m := make(map[string]int64)
	sampleLen := len(vector)
	for i, sample := range vector {
		klog.Infof("Sample %d/%d: %s", i+1, sampleLen, sample)
		nodeName := string(sample.Metric["host"])
		latency := int64(sample.Value)
		klog.Infof("Extracted latency value %d for node '%s'", latency, nodeName)
		m[nodeName] = latency
	}

	return m
}
