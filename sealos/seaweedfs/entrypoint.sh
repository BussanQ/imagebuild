#!/usr/bin/env bash
set -e
cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1

NAME=${NAME:-"seaweedfs"}
NAMESPACE=${NAMESPACE:-"seaweedfs"}
CHARTS=${CHARTS:-"./charts/seaweedfs"}
HELM_OPTS=${HELM_OPTS:-""}

helm upgrade -i ${NAME} ${CHARTS} -n ${NAMESPACE} --create-namespace ${HELM_OPTS}
