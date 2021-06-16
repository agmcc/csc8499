#!/bin/bash
source vpn_config.sh
envsubst < 40-route.tmpl > 40-route.tmp
sudo mv 40-route.tmp /lib/dhcpcd/dhcpcd-hooks/40-route
echo "Reboot for changes to be applied"
