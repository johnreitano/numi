#!/usr/bin/env bash

set -e
set -x

TLS_CERTIFICATE_EMAIL=$1
RPC_HOST=$2
API_HOST=$3

# stop service
sudo systemctl stop numi.service || :

# download latest numid from github
rm -rf numid-download
mkdir numid-download
cd numid-download
DOWNLOAD_URL=$(curl -s https://api.github.com/repos/johnreitano/numi/releases/latest | jq -r '.assets[] | select(.name|match("linux_amd64.tar.gz$")) | .browser_download_url')
wget -q ${DOWNLOAD_URL}
tar xzf *linux_amd64.tar.gz
sudo mv numid /usr/local/bin
cd ..

# move uploaded home dir into position
rm -rf upload/
tar xzf upload.tgz
cp -r upload/numi-home /home/ubuntu/.numi

# configure additional items
upload/install-generic-cert.sh ${TLS_CERTIFICATE_EMAIL} ${RPC_HOST}
upload/install-nginx-cert.sh ${TLS_CERTIFICATE_EMAIL} ${API_HOST} 1317
sudo cp upload/numi.service /etc/systemd/system/numi.service
sudo chmod 664 /etc/systemd/system/numi.service
sudo systemctl daemon-reload

# start numid service
sudo systemctl start numi.service
sleep 1
sudo systemctl status -l numi.service --no-pager
sudo systemctl enable numi.service
