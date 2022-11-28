#!/usr/bin/env bash

set -e
# set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

# download latest mac version of client
rm -rf deploy/download
mkdir deploy/download
pushd deploy/download
if [[ $(uname -p) = "arm" ]]; then GOARCH=arm64; else GOARCH=amd64; fi
DOWNLOAD_URL=$(curl -s https://api.github.com/repos/johnreitano/numi/releases/latest | jq -r '.assets[] | select(.name|match("darwin_'${GOARCH}'.tar.gz")) | .browser_download_url')
wget -q ${DOWNLOAD_URL}
tar xzf *.tar.gz
popd
NUMID=deploy/download/numid

rm -rf ~/.numi
${NUMID} init --chain-id numi-local myclient >/dev/null 2>/dev/null
deploy/add-test-keys.sh ~/.numi >/dev/null 2>/dev/null

echo "Users in numi blockchain before adding user:"
${NUMID} query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'

yes | ${NUMID} keys delete newbie --keyring-backend test --home ~/.numi >/dev/null 2>/dev/null || :
${NUMID} keys add newbie --keyring-backend test --home ~/.numi >/dev/null 2>/dev/null
NEWBIE_ADDR=$(${NUMID} keys show newbie -a --keyring-backend test --home ~/.numi)

echo -e "\nPlease enter info for new user:\n"
echo -n "First name: "; read FIRST_NAME
echo -n "Last name: "; read LAST_NAME
echo -n "Bio: "; read BIO

echo -e "\nIn another terminal window please enter the following command:\n"

BOB_ADDR=$(${NUMID} keys show bob -a --keyring-backend test --home ~/.numi)
OLIVER_ADDR=$(${NUMID} keys show oliver -a --keyring-backend test --home ~/.numi)
USER_ID=$(uuidgen)

echo ${NUMID} tx numi create-and-verify-user ${USER_ID} \"${FIRST_NAME}\" \"${LAST_NAME}\" USA California \"San Diego\" \"$BIO\" ${BOB_ADDR} ${NEWBIE_ADDR} --keyring-backend test --from ${OLIVER_ADDR} --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657

sleep 7

echo -e -n "\nPress any key once you have run the command above..."; read DUMMY

echo "\nUsers in numi blockchain after adding user:"
${NUMID} query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'
