#!/usr/bin/env bash

set -e
set -x

TLS_CERTIFICATE_EMAIL=$1
DOMAIN_PREFIX=$2
DNS_ZONE_NAME=$3

rm -rf upload/
tar xzf upload.tgz

sudo systemctl stop numi.service || :

sudo cp upload/numid /usr/local/bin/
cp -r upload/numi-home /home/ubuntu/.numi

upload/install-generic-cert.sh ${TLS_CERTIFICATE_EMAIL} ${DOMAIN_PREFIX}validator-${i}-rpc.${DNS_ZONE_NAME}
upload/install-nginx-cert.sh ${TLS_CERTIFICATE_EMAIL} ${DOMAIN_PREFIX}validator-${i}-api.${DNS_ZONE_NAME} 1317

sudo cp upload/numi.service /etc/systemd/system/numi.service
sudo chmod 664 /etc/systemd/system/numi.service
sudo systemctl daemon-reload
sudo systemctl start numi.service
sleep 1
sudo systemctl status -l numi.service --no-pager
sudo systemctl enable numi.service
