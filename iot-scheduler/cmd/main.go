package main

import (
	"iot-scheduler/pkg/dummy"

	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	klog.Info("Hello, scheduluer")

	command := app.NewSchedulerCommand(
		app.WithPlugin(dummy.Name, dummy.New),
	)

	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}
