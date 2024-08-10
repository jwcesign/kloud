#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

controller-gen crd paths=./pkg/apis/cluster/... output:crd:dir=./config/crds
