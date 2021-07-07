package metrics

import (
	"k8s.io/klog/v2"
)

// TODO: Stubbed
func GetLatencies(nodes []string) map[string]int64 {
	klog.Infof("Retrieving latency metrics for nodes: %s", nodes)

	latencies := map[string]int64{
		"instance-2": 45,
		"instance-3": 30,
		"pilab-01":   400,
		"pilab-05":   200,
		"pilab-06":   150,
	}

	m := make(map[string]int64)

	for _, n := range nodes {
		l := latencies[n]
		m[n] = l
		klog.Infof("%s: %d", n, l)
	}

	return m
}
