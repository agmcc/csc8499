#!/bin/bash
master=$(sudo kubectl get nodes | grep master | awk '{print$1}')
sudo kubectl taint nodes ${master} node-role.kubernetes.io/master=true:NoSchedule

