package dummy

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type Dummy struct {
	handle framework.Handle
}

var _ = framework.ScorePlugin(&Dummy{})

const Name = "Dummy"

func (ps *Dummy) Name() string {
	fmt.Println("Getting name")
	return Name
}

func (ps *Dummy) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Info("Calculating score!")
	fmt.Printf("Score for node: %s\n", nodeName)
	return 1, nil
}

func (ps *Dummy) ScoreExtensions() framework.ScoreExtensions {
	klog.Info("Score extensions!")
	fmt.Println("Score extensions (fmt)")
	return ps
}

func (ps *Dummy) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	klog.Info("NormalizeScore!")
	fmt.Println("NormalizeScore (fmt)")
	return nil
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	klog.Info("Creating dummy plugin")
	fmt.Println("New...")
	return &Dummy{handle: h}, nil
}
