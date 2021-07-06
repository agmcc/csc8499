package main

import (
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	klog.Info("Hello, scheduluer")

	command := app.NewSchedulerCommand()

	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}
