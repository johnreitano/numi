#!/usr/bin/env bash

set -e
# set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

VALIDATOR_IPS=($*)

# create and initialize the the validator config dirs
mkdir -p deploy/node-config
rm -rf deploy/node-config/*

VALIDATOR_MNEMONICS=("gun quick banner word mutual pet sort run illness behind pull stock crazy talk actor icon help gym young census decorate swamp two plunge"
 "mule multiply combine frown aim window top weekend frown cancel turn token canoe thumb attitude flame execute purpose chest design winner enable coconut retire" "business bless fuel joy lady volcano odor tribe virus have effort rate mouse disease general view mention evoke lend expect frozen trend shrimp flavor")
for i in ${!VALIDATOR_IPS[@]}; do 
    MONIKER=validator${i}
    HOME_DIR=deploy/node-config/validator${i}
    mkdir deploy/node-config/${MONIKER}
    echo ${VALIDATOR_MNEMONICS[$i]} | numid init --chain-id numi-testnet-1 --recover --home ${HOME_DIR} ${MONIKER}
done

# set persistent peers in variable
P2P_PERSISTENT_PEERS=""
for i in ${!VALIDATOR_IPS[@]}; do 
    HOME_DIR=deploy/node-config/validator${i}
    VALIDATOR_NODE_ID=$(numid tendermint show-node-id --home ${HOME_DIR})
    P2P_PERSISTENT_PEERS="${P2P_PERSISTENT_PEERS}${VALIDATOR_NODE_ID}@${VALIDATOR_IPS[$i]}:26656,"
done

# update config files with external_address, persistent_peers and other values
for i in ${!VALIDATOR_IPS[@]}; do
  MONIKER=validator${i}
  HOME_DIR=deploy/node-config/${MONIKER}

  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.external_address" "tcp://${VALIDATOR_IPS[$i]}:26656"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".p2p.persistent_peers" "${P2P_PERSISTENT_PEERS}"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_cert_file" "/home/ubuntu/cert/fullchain.pem"
  dasel put string -f ${HOME_DIR}/config/config.toml -p toml ".rpc.tls_key_file" "/home/ubuntu/cert/privkey.pem"
  dasel put bool -f ${HOME_DIR}/config/app.toml -p toml ".api.enable" true
  dasel put string -f ${HOME_DIR}/config/app.toml -p toml ".api.address" "tcp://localhost:1317"
done

# generate genesis transaction on each validator; for secondary validators, copy genesis transaction to primary validator
for i in ${!VALIDATOR_IPS[@]}; do
  MONIKER=validator${i}
  HOME_DIR=deploy/node-config/${MONIKER}
  yes | numid keys delete ${MONIKER}-key --keyring-backend test --home ${HOME_DIR} 2>/dev/null || :
  echo ${VALIDATOR_MNEMONICS[$i]} | numid keys add ${MONIKER}-key --keyring-backend test --home ${HOME_DIR} --recover
  ADDR=$(numid keys show ${MONIKER}-key -a --keyring-backend test --home ${HOME_DIR})
  echo ${MONIKER}: $ADDR
  numid add-genesis-account --keyring-backend test --home ${HOME_DIR} ${ADDR} 2000000000unumi || :
  numid gentx ${MONIKER}-key 1000000000unumi --chain-id numi-testnet-1 --moniker=${MONIKER} --keyring-backend test --home ${HOME_DIR}

  if [[ ${i} != "0" ]]; then
    cp ${HOME_DIR}/config/gentx/* deploy/node-config/validator0/config/gentx/    
    numid add-genesis-account --keyring-backend test --home deploy/node-config/validator0 ${ADDR} 2000000000unumi || :
  fi
done

# collect genesis transactions on primary node
MONIKER=validator0
HOME_DIR=deploy/node-config/${MONIKER}
numid collect-gentxs --home ${HOME_DIR}

# calculate oracle addresses using menmonics (not safe for prod!)
ORACLE_MNEMONICS=("engage unhappy soft business govern transfer spider buzz soda boost robot ugly fix suggest source key sell silk shaft online enforce economy capable news" "cannon problem manual elder shop hero enable walnut exclude hour sand connect tower puppy frown mean ten member grace tower phone shop civil february")
ORACLE_ADDRS=()
PRIMARY_HOME_DIR=deploy/node-config/validator0
for i in ${!ORACLE_MNEMONICS[@]}; do 
  yes | numid keys delete temp-key --keyring-backend test --home ${PRIMARY_HOME_DIR} 2>/dev/null || :
  echo ${ORACLE_MNEMONICS[$i]} | numid keys add temp-key --keyring-backend test --home ${PRIMARY_HOME_DIR} --recover
  ORACLE_ADDRS+=($(numid keys show temp-key -a --keyring-backend test --home ${PRIMARY_HOME_DIR}))
  yes | numid keys delete temp-key --keyring-backend test --home ${PRIMARY_HOME_DIR} 2>/dev/null || :
done

# update module parameters within genesis file on primary node
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.crisis.constant_fee.denom" "numi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.gov.deposit_params.min_deposit[0].denom" "numi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.gov.voting_params.voting_period" "60s"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.mint.params.mint_denom" "numi"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.staking.params.bond_denom" "numi"
echo dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.numi.params.identityVerifiers" "${ORACLE_ADDRS[0]},${ORACLE_ADDRS[1]}"
dasel put string -f ${PRIMARY_HOME_DIR}/config/genesis.json -p json ".app_state.numi.params.identityVerifiers" "numi1wdnwe6tt2wz3glwe2d0cjmwk8nzvracwp6p8c7,numi13crpqdukn5l3gr4pzzcjzcl6fpx7rhay8uvy44"

# copy genesis file from primary node to secondary nodes
for i in ${!VALIDATOR_IPS[@]}; do
  if [[ "${i}" != "0" ]]; then
    MONIKER=validator${i}
    HOME_DIR=deploy/node-config/${MONIKER}
    cp ${PRIMARY_HOME_DIR}/config/genesis.json ${HOME_DIR}/config/
  fi
done

