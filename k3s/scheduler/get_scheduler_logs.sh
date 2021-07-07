#!/bin/bash
pod=$(sudo kubectl get pods -n kube-system | grep iot-scheduler | awk '{print $1}')

if [ -z "${pod}" ]
then
    echo "Scheduler pod not found - exiting"
    exit 1
fi

echo "Showing logs for scheduler pod: ${pod}"

sudo kubectl logs -n kube-system $pod -f
