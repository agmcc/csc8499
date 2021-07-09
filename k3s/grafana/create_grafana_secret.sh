#!/bin/bash
read -sp "Enter Grafana admin password: " password

sudo kubectl create secret generic \
grafana-admin-secret \
--from-literal=password=$password \
-n monitoring
