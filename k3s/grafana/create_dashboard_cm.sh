#!/bin/bash
sudo kubectl create configmap \
grafana-dashboard-conf \
--from-file=./dashboard/metrics-dashboard.json \
-n monitoring
