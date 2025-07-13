#!/usr/bin/env bash

set -e
cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1

NAME=${NAME:-"fluid"}
NAMESPACE=${NAMESPACE:-"fluid"}
CHARTS="./charts/fluid"
HELM_OPTS=${HELM_OPTS:-""}

function install(){
  helm upgrade -i ${NAME} ${CHARTS} -n ${NAMESPACE} --create-namespace \
  --set csi.config.hostNetwork=true ${HELM_OPTS}
}

install
