#!/bin/bash

set -e

SCRIPT_PATH="$(cd $(dirname "$0") >/dev/null 2>&1 && pwd)"
PROJECT_PATH="${SCRIPT_PATH}/../"
source ${PROJECT_PATH}/docker/SETTINGS



# -----------------------------------------------------------------------------
push_img() {
   _log "Push images"
   docker push "$LIBRARY"/"$IMAGENAME":"$IMGVERS1"
   docker push "$LIBRARY"/"$IMAGENAME":"$IMGVERS2"
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

_log "--- Detecting tools"
find_tool realpath || FIND_RES=1
find_tool docker || FIND_RES=1
find_tool basename || FIND_RES=1
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

push_img
