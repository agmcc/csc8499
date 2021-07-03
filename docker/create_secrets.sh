#!/bin/bash
read -sp "Enter Grafana admin password: " grafana_admin_password
echo -n $grafana_admin_password | docker secret create grafana_admin_password -
