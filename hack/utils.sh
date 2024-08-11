#!/usr/bin/env bash

function utils::create_gopath_tree() {
  # $1: the root path of repo
  local repo_root=$1
  # $2: go path
  local go_path=$2

  local cloudpilot_go_package="github.com/jwcesign/kloud"

  local go_pkg_dir="${go_path}/src/${cloudpilot_go_package}"
  go_pkg_dir=$(dirname "${go_pkg_dir}")

  mkdir -p "${go_pkg_dir}"

  if [[ ! -e "${go_pkg_dir}" || "$(readlink "${go_pkg_dir}")" != "${repo_root}" ]]; then
    ln -snf "${repo_root}" "${go_pkg_dir}"
  fi
}
