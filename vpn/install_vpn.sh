#!/bin/bash
# Run as root
if [ "$EUID" -ne 0 ]
  then echo "Script must be run as root"
  exit 1
fi

# Install
apt update
apt install -y strongswan

# Config
cp gcloud.conf /etc
conf="include /etc/gcloud.conf"
confFile=/etc/ipsec.conf
grep -qF $conf $confFile || echo $conf >> $confFile

# Secret
cp gcloud.secrets /etc
read -p "Enter PSK: " psk
sed -i "s/<PSK>/${psk}/" /etc/gcloud.secrets
secret="include /etc/gcloud.secrets"
secretsFile=/etc/ipsec.secrets
grep -qF $secret $secretsFile || echo $secret >> $secretsFile

# Restart
service strongswan restart