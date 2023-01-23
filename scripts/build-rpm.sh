#!/bin/bash

trap do_exit INT
trap do_exit EXIT

SCRIPT_PATH="$(cd $(dirname "$0") >/dev/null 2>&1 && pwd)"
PROJECT_PATH="${SCRIPT_PATH}/../"

DOCKER_BASE_IMAGE="centos:8"
DOCKER_CONTAINER_NAME="peoplevine-auth-rpm-build_${JOB_NAME:-NONE}_${BUILD_NUMBER:-0}"
DOCKER_USER=$([ $(id -u) == '0' ] && echo "root" || echo "jenkins")

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
docker_rm() {
    _log "Remove container [${DOCKER_CONTAINER_NAME}]"
    TEMP_LOG=$(sudo docker rm -f ${DOCKER_CONTAINER_NAME} 2>&1)
}

# -----------------------------------------------------------------------------
docker_shell() {
    _log "Exec shell"
    sudo docker exec -ti -u ${DOCKER_USER} ${DOCKER_CONTAINER_NAME} /bin/bash
}

# -----------------------------------------------------------------------------
docker_exec() {
     local PARAM_COMMAND="${@}"
    _log "Exec command [${PARAM_COMMAND}]"
    sudo docker exec --tty -u ${DOCKER_USER} ${DOCKER_CONTAINER_NAME} /bin/bash -c ". /etc/bashrc; ${PARAM_COMMAND}"
    RET_RES="${?}"
    if [ "${RET_RES}" != "0" ]; then
       _log "ERROR: execute command [${PARAM_COMMAND}]"
       return 127
    else
       return 0
    fi
}

# -----------------------------------------------------------------------------
docker_run() {
    _log "Run container"

    local PARAM_DOCKER_PARAMETERS="${@}"
    docker_rm
    sudo docker pull ${DOCKER_BASE_IMAGE}
    sudo docker run -d --rm --name "${DOCKER_CONTAINER_NAME}" \
    --privileged --sysctl=net.ipv6.conf.all.disable_ipv6=0 \
    -e BUILD_UID="$(id -u)" -e BUILD_GID="$(id -g)" \
    -e BUILD_NUMBER="${BUILD_NUMBER}" -e CI_PIPELINE_ID="${CI_PIPELINE_ID}" \
    -v ${PROJECT_PATH}/build:/build \
    ${DOCKER_BASE_IMAGE} bash -c "while true; do sleep 10; done"
    if [ "${?}" != "0" ]; then
       _log "ERROR: Docker run error"
       return 127
    fi
   _log "OK"
#   return 0
}

# -----------------------------------------------------------------------------
# -----------------------------------------------------------------------------
# -----------------------------------------------------------------------------
do_exit() {
  docker_rm
  exit ${1}
}

# -----------------------------------------------------------------------------
do_prep_rpm_src() {
  _log "Prepare RPM source"
  rm -rf ${PROJECT_PATH}/build
  mkdir -p ${PROJECT_PATH}/build/rpm/peoplevine-auth
  cp -rf ${PROJECT_PATH}/pkg ${PROJECT_PATH}/build/rpm/peoplevine-auth
  cp -rf ${PROJECT_PATH}/peoplevine-auth ${PROJECT_PATH}/build/rpm/peoplevine-auth
  cp -rf ${PROJECT_PATH}/systemd ${PROJECT_PATH}/build/rpm/peoplevine-auth
  cp -rf ${PROJECT_PATH}/peoplevine-auth.spec ${PROJECT_PATH}/build/rpm/
  cp -rf ${PROJECT_PATH}/README.md ${PROJECT_PATH}/build/rpm/peoplevine-auth
  pushd ${PROJECT_PATH}/build/rpm
    tar czf peoplevine-auth.tgz peoplevine-auth
  popd
}

# -----------------------------------------------------------------------------

_log "--- Detecting tools"
find_tool realpath || FIND_RES=1
find_tool docker || FIND_RES=1
find_tool basename || FIND_RES=1
find_tool rsync || FIND_RES=1
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

do_prep_rpm_src

docker_run
docker_exec "dnf -y --disablerepo '*' --enablerepo=extras swap centos-linux-repos centos-stream-repos"
docker_exec "dnf -y install epel-release"
docker_exec "dnf -y install mock"
docker_exec "mock --buildsrpm --spec /build/rpm/peoplevine-auth.spec --sources /build/rpm  --resultdir=/build/rpm"
docker_exec "mock --resultdir=/build/rpm /build/rpm/*.rpm && cp -f /build/rpm/*.x86_64.rpm /build && rm -rf /build/rpm"
 