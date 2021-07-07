package plugin

import (
	"context"
	"math"

	"iot-scheduler/pkg/metrics"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type LatencyAware struct {
	handle framework.Handle
}

type LatencyAwareState struct {
	m map[string]int64
}

const (
	Name                 = "LatencyAware"
	LatencyAwareStateKey = "LatencyAwareState"
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
	latencies := metrics.GetLatencies(nodeNames)
	m := make(map[string]int64)
	for k, v := range latencies {
		m[k] = math.MaxInt64 - v
	}
	latencyAwareState := &LatencyAwareState{m}
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
	latency := latencyAwareState.m[nodeName]
	klog.Infof("Loaded latency value for node %s: %d", nodeName, latency)
	return latency, nil
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
