#!/usr/bin/env bash

set -e
set -x

SCRIPT_DIR=$(dirname $(readlink -f $0))
cd ${SCRIPT_DIR}/..

HOME_DIR=$1

# add keys for oracle account oliver and olivia and test accounts alice, bob and carol
USERS=(
  "oliver;engage unhappy soft business govern transfer spider buzz soda boost robot ugly fix suggest source key sell silk shaft online enforce economy capable news"
  "olivia;cannon problem manual elder shop hero enable walnut exclude hour sand connect tower puppy frown mean ten member grace tower phone shop civil february"
  "alice;grant desk armor salmon mixture grid amateur auto timber crowd honey elder scissors radar smile mutual cheap stadium diesel sound design weird brisk join"
  "bob;inject fold fluid champion doctor figure since waste pig similar nation benefit wrong picture during finger sister guilt chat sight avocado lottery risk citizen"
  "carol;rookie pudding garage over reform use paper speed runway tell speed client hazard movie table neutral damage episode keep topple brain zone gap mobile"
)
for i in ${!USERS[@]}; do
  IFS=';' read -ra USER <<< "${USERS[$i]}"
  NAME=${USER[0]}
  MNEMONIC=${USER[1]}
  yes | numid keys delete ${NAME} --keyring-backend test --home ${HOME_DIR} 2>/dev/null || :
  echo ${MNEMONIC} | numid keys add ${NAME} --keyring-backend test --home ${HOME_DIR} --recover
  ADDR=$(numid keys show ${NAME} -a --keyring-backend test --home ${HOME_DIR})
  echo "Added account for key ${NAME} with ${ADDR}"
done
