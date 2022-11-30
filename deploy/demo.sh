#!/usr/bin/env bash

set -e
# set -x

SCRIPT_DIR=$(dirname $0)
cd ${SCRIPT_DIR}/..

# download latest mac version of client
rm -rf deploy/download
mkdir deploy/download
pushd deploy/download >/dev/null
if [[ $(uname -p) = "arm" ]]; then GOARCH=arm64; else GOARCH=amd64; fi
DOWNLOAD_URL=$(curl -s https://api.github.com/repos/johnreitano/numi/releases/latest | jq -r '.assets[] | select(.name|match("darwin_'${GOARCH}'.tar.gz")) | .browser_download_url')
wget -q ${DOWNLOAD_URL}
tar xzf *.tar.gz
popd >/dev/null
NUMID=deploy/download/numid

rm -rf ~/.numi
${NUMID} init --chain-id numi-local myclient >/dev/null 2>/dev/null
deploy/add-test-keys.sh ~/.numi >/dev/null 2>/dev/null

printf "\nCOMMAND 1: Show list of registered users BEFORE adding a user\n"

GREEN='\033[0;32m'
NO_COLOR='\033[0m'

printf "\nCopy the following command and paste it into the right terminal:\n"

printf "\n${GREEN}${NUMID} query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'${NO_COLOR}\n"

printf "\nPress any key when done..."; read DUMMY

yes | ${NUMID} keys delete newbie --keyring-backend test --home ~/.numi >/dev/null 2>/dev/null || :
${NUMID} keys add newbie --keyring-backend test --home ~/.numi >/dev/null 2>/dev/null
NEWBIE_ADDR=$(${NUMID} keys show newbie -a --keyring-backend test --home ~/.numi)


printf "\nPlease enter info for new user:\n"
printf "\nFirst name: "; read FIRST_NAME
printf "\nLast name: "; read LAST_NAME
printf "\nBio: "; read BIO

printf "\nCOMMAND 2: Register a new user\n"

BOB_ADDR=$(${NUMID} keys show bob -a --keyring-backend test --home ~/.numi)
OLIVER_ADDR=$(${NUMID} keys show oliver -a --keyring-backend test --home ~/.numi)
USER_ID=$(uuidgen)

printf "\nCopy the following command and paste it into the right terminal:\n"

printf "\n${GREEN}${NUMID} tx numi create-and-verify-user ${USER_ID} \"${FIRST_NAME}\" \"${LAST_NAME}\" USA California \"San Diego\" \"$BIO\" ${BOB_ADDR} ${NEWBIE_ADDR} --keyring-backend test --from ${OLIVER_ADDR} --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657${NO_COLOR}\n"

printf "\nPress any key when done..."; read DUMMY

sleep 4

printf "\nCOMMAND 3: Show list of registered users AFTER adding a user\n"

printf "\nCopy the following command and paste it into the right terminal:\n"

printf "\n${GREEN}${NUMID} query numi list-user --chain-id numi-testnet-1 --node https://testnet-seed-0-rpc.numi.oktryme.com:26657 --output json | jq '.user'${NO_COLOR}\n"

printf "\nPress any key when done..."; read DUMMY

printf "\nALL DONE\n"
