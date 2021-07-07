package plugin

import (
	"context"

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
	m := make(map[string]int64)
	for _, n := range nodes {
		klog.Info("Pre-scoring node '%s' based on name length", n)
		nodeName := n.Name
		length := int64(len(nodeName))
		klog.Info("Length: %s", length)
		m[nodeName] = length
	}
	latencyAwareState := &LatencyAwareState{m}
	state.Write(LatencyAwareStateKey, latencyAwareState)
	return nil
}

func (ps *LatencyAware) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Info("Calculating score for node: %s", nodeName)
	c, err := state.Read(LatencyAwareStateKey)
	if err != nil {
		klog.Fatal("Unable to read state: %s", err)
	}
	latencyAwareState, ok := c.(*LatencyAwareState)
	if !ok {
		klog.Fatal("Unable to convert cycle state to latency state")
	}
	latency := latencyAwareState.m[nodeName]
	klog.Info("Loaded latency value for node %s: %s", latency, nodeName)
	return latency, nil
}

func (ps *LatencyAware) ScoreExtensions() framework.ScoreExtensions {
	return ps
}

func (ps *LatencyAware) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.Info("Normalizing score")
	klog.Info("Pre-normalized scores: %s", scores)
	highest := int64(0)
	for _, s := range scores {
		if s.Score > highest {
			highest = s.Score
		}
	}
	for _, s := range scores {
		s.Score = s.Score * framework.MaxNodeScore / highest
	}
	klog.Info("Normalized scores: %s", scores)
	return nil
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	klog.Info("Creating %s plugin", Name)
	return &LatencyAware{handle: h}, nil
}
