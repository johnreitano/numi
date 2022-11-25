#!/usr/bin/env bash

set -e
set -x

TLS_CERTIFICATE_EMAIL=$1
RPC_DOMAIN=$2
API_DOMAIN=$3
API_PORT=$4

# stop service
sudo systemctl stop numi.service || :
sleep 1

# update packages
sudo apt update -y
sudo snap install core
sudo snap refresh core
if [[ -z "$(which jq)" ]]; then
    sudo apt install -y jq
fi
if [[ -z "$(which dasel)" ]]; then
    sudo wget -qO /usr/local/bin/dasel https://github.com/TomWright/dasel/releases/latest/download/dasel_linux_amd64
    sudo chmod a+x /usr/local/bin/dasel
fi
sudo apt remove -y certbot
if [[ -z "$(which certbot)" ]]; then
    sudo snap install --classic certbot
    sudo ln -s /snap/bin/certbot /usr/bin/certbot
fi

# set maximum number of open files to 4096 -- needed by numid
ulimit -n 4096 

# download latest numid from github
rm -rf numid-download
mkdir numid-download
cd numid-download
DOWNLOAD_URL=$(curl -s https://api.github.com/repos/johnreitano/numi/releases/latest | jq -r '.assets[] | select(.name|match("linux_amd64.tar.gz$")) | .browser_download_url')
wget -q ${DOWNLOAD_URL}
tar xzf *linux_amd64.tar.gz
sudo mv numid /usr/local/bin
cd ..

# move uploaded home dir into default location
rm -rf upload/
tar xzf upload.tgz
rm -rf .numi
cp -r upload/numi-home .numi

# install certs
upload/install-certs.sh ${TLS_CERTIFICATE_EMAIL} ${RPC_DOMAIN} ${API_DOMAIN} ${API_PORT}

# configure and start numi.service
sudo cp upload/numi.service /etc/systemd/system/numi.service
sudo chmod 664 /etc/systemd/system/numi.service
sudo systemctl daemon-reload
sudo systemctl start numi.service
sleep 1
sudo systemctl status -l numi.service --no-pager
sudo systemctl enable numi.service
