#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# For all commands, the working directory is the parent directory(repo root).
REPO_ROOT=$(pwd)
cd "${REPO_ROOT}"

export GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')
export PATH=$PATH:$GOPATH/bin

boilerplate="${REPO_ROOT}"/hack/boilerplate/boilerplate.go.txt

go_path="${REPO_ROOT}/_go"
cleanup() {
  rm -rf "${go_path}"
}
trap "cleanup" EXIT SIGINT

cleanup

source "${REPO_ROOT}"/hack/utils.sh
utils::create_gopath_tree "${REPO_ROOT}" "${go_path}"
export GOPATH="${go_path}"

echo "Generating with deepcopy-gen"
deepcopy-gen \
  --go-header-file "${boilerplate}" \
  --output-file-base zz_generated.deepcopy \
  --input-dirs github.com/jwcesign/kloud/pkg/apis/clustermigration/v1alpha1

echo "Generating with register-gen"
register-gen \
  --go-header-file "${boilerplate}" \
  --output-file-base zz_generated.register \
  --input-dirs github.com/jwcesign/kloud/pkg/apis/clustermigration/v1alpha1

echo "Generating with conversion-gen"
conversion-gen \
  -O zz_generated.conversion \
  --go-header-file "${boilerplate}" \
  --input-dirs github.com/jwcesign/kloud/pkg/apis/clustermigration/v1alpha1


