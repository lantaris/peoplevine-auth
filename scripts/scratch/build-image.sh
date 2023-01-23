#!/bin/bash

set -e

SCRIPT_PATH="$(cd $(dirname "$0") >/dev/null 2>&1 && pwd)"
PROJECT_PATH="${SCRIPT_PATH}/../"
source ${PROJECT_PATH}/docker/SETTINGS

# -----------------------------------------------------------------------------
crate_dir() {
   rm -rf ${PROJECT_PATH}/build
   mkdir -p ${PROJECT_PATH}/build/go
}

# -----------------------------------------------------------------------------
build_img() {
   _log "Build docker image [$LIBRARY/$IMAGENAME:$IMGVERS1]"
   export BUILDDATE=$(date +%Y%m%d)
   export IMAGENAME
   export IMGVERS1
   export IMGVERS2
   j2 -f env -o ${PROJECT_PATH}/build/Dockerfile     ${PROJECT_PATH}/docker/Dockerfile.tmpl
   pushd ${PROJECT_PATH}/build
     docker build -t "${LIBRARY}"/"${IMAGENAME}":"${IMGVERS1}" -t "${LIBRARY}"/"${IMAGENAME}":"${IMGVERS2}" .
   popd

}

# -----------------------------------------------------------------------------
_log() {
  echo -e "--- $@"
}

# -----------------------------------------------------------------------------
find_tool() {
    local PARAM_CMD="${1}"
    RES_TMP="$(which ${PARAM_CMD} 2>&1)"
    if [ "${?}" != "0" ]; then
      echo "Tools [${PARAM_CMD}] NOT FOUND (Please install first)"
      return 127
    fi
    return 0
}

# -----------------------------------------------------------------------------
docker_check_privileges() {
    sudo docker ps 2>&1 >/dev/null
    if [ "${?}" != "0" ]; then
       _log "ERROR: Docker not runing or not accessed"
       return 127
    else
       _log "OK"
       return 0
    fi
}

# -----------------------------------------------------------------------------
build_go() {
  cp -rf ${PROJECT_PATH}/pkg ${PROJECT_PATH}/build/go
  cp -rf ${PROJECT_PATH}/peoplevine-auth ${PROJECT_PATH}/build/go
  pushd ${PROJECT_PATH}/build/go
    export GOPATH=$(pwd)
    cd peoplevine-auth
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ${PROJECT_PATH}/build/peoplevine-auth
  popd
  rm -rf ${PROJECT_PATH}/build/go
}


# -----------------------------------------------------------------------------

_log "--- Detecting tools"
find_tool realpath || FIND_RES=1
find_tool docker || FIND_RES=1
find_tool basename || FIND_RES=1
find_tool rsync || FIND_RES=1
find_tool j2 || FIND_RES=1
if [ "${FIND_RES}" == "1" ]; then
 _log "Please install required tools"
 exit 127
else
 _log "OK".
fi

# Check docker access
_log "--- Check Docker privileges"
docker_check_privileges
if [ "${?}" != "0" ]; then
  exit 127
fi

crate_dir

build_go

build_img
 