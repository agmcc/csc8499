package plugin

import (
	"context"
	"math"

	"iot-scheduler/pkg/metrics"
	"iot-scheduler/pkg/utils"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type LatencyAware struct {
	handle framework.Handle
}

type LatencyAwareState struct {
	metrics         map[string]NodeMetrics
	maxLatency      int64
	maxResponseTime int64
}

type NodeMetrics struct {
	latency      int64
	responseTime int64
}

const (
	Name                 = "LatencyAware"
	LatencyAwareStateKey = "LatencyAwareState"
	MetricsAddress       = "http://prometheus-service.monitoring.svc.cluster.local:9090"
)

var _ = framework.PreScorePlugin(&LatencyAware{})
var _ = framework.ScorePlugin(&LatencyAware{})

func (s *LatencyAwareState) Clone() framework.StateData {
	return s
}

func (ps *LatencyAware) Name() string {
	return Name
}

func (ps *LatencyAware) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	klog.Info("Entered PreScore")

	var nodeNames []string
	for _, n := range nodes {
		nodeNames = append(nodeNames, n.Name)
	}

	api := metrics.NewAPI(MetricsAddress)

	latencies := metrics.GetNetworkLatencies(ctx, api, nodeNames)
	klog.Info("Retreived latency metrics: ", latencies)

	responseTimes := metrics.GetResponseTimes(ctx, api, nodeNames)
	klog.Info("Retrieved response time metrics: ", responseTimes)

	filledLatencies := utils.FillMissingValues(latencies, metrics.Missing)
	klog.Info("Filled latency metrics: ", filledLatencies)

	filledResponseTimes := utils.FillMissingValues(responseTimes, metrics.Missing)
	klog.Info("Filled response time metrics: ", filledResponseTimes)

	m := make(map[string]NodeMetrics)
	var maxLatency int64 = -math.MaxInt64
	var maxResponseTime int64 = -math.MaxInt64

	for _, n := range nodeNames {
		l := filledLatencies[n]
		if l > maxLatency {
			maxLatency = l
		}
		r := filledResponseTimes[n]
		if r > maxResponseTime {
			maxResponseTime = r
		}
		m[n] = NodeMetrics{latency: l, responseTime: r}
	}
	klog.Infof("Max latency is %d, max response time is %d", maxLatency, maxResponseTime)

	latencyAwareState := &LatencyAwareState{metrics: m, maxLatency: maxLatency, maxResponseTime: maxResponseTime}
	klog.Infof("LatencyAwareState: %v", latencyAwareState)
	state.Write(LatencyAwareStateKey, latencyAwareState)

	return nil
}

func (ps *LatencyAware) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Infof("Calculating score for node: %s", nodeName)
	c, err := state.Read(LatencyAwareStateKey)
	if err != nil {
		klog.Fatal("Unable to read state", err)
	}
	latencyAwareState, ok := c.(*LatencyAwareState)
	if !ok {
		klog.Fatal("Unable to convert cycle state to latency state")
	}

	metrics := latencyAwareState.metrics[nodeName]
	latencyScore := latencyAwareState.maxLatency - metrics.latency
	responseTimeScore := latencyAwareState.maxResponseTime - metrics.responseTime

	score := latencyScore + responseTimeScore
	klog.Infof("Combined score for node %s: %d (latency score %d, response time score %d)", nodeName, score, latencyScore, responseTimeScore)
	return score, nil
}

func (ps *LatencyAware) ScoreExtensions() framework.ScoreExtensions {
	return ps
}

func (ps *LatencyAware) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.Info("Normalizing score")
	var highest int64 = -math.MaxInt64
	highestScoreNode := ""
	for _, s := range scores {
		if s.Score > highest {
			highest = s.Score
			highestScoreNode = s.Name
		}
	}
	klog.Infof("Node with highest score is %s: %d", highestScoreNode, highest)
	for i, s := range scores {
		normalized := s.Score * framework.MaxNodeScore / highest
		klog.Infof("Normalized score for %s: %d (was %d)", s.Name, normalized, s.Score)
		scores[i].Score = normalized
	}
	return nil
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	klog.Infof("Creating %s plugin", Name)
	return &LatencyAware{handle: h}, nil
}
