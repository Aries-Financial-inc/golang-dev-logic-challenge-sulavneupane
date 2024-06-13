#!/usr/bin/env bash

set -o pipefail

export COLOR_RED="\033[31m"
export COLOR_NORMAL="\033[39m"
export COLOR_ERROR="${COLOR_RED}"

error (){
  echo -e "${COLOR_ERROR}ERROR:${COLOR_NORMAL} $*" 1>&2
}


# git directory root
ROOT_DIR="$(git rev-parse --show-toplevel)"

# grab major version from go.mod
GOMOD_VER="$(awk '/go [0-9].[0-9]+/{print $2; exit}' "$ROOT_DIR/go.mod")"
# grab the major minor version of go
LOCAL_VER="$(go version | awk '{gsub(/go/, ""); print $2}' | cut -d'.' -f1,2)"

if [ "$GOMOD_VER" != "$LOCAL_VER" ]; then
  error "go version mismatch!"
  echo "PROJECT: $GOMOD_VER"
  echo "LOCAL:   $LOCAL_VER"
  exit 1
fi