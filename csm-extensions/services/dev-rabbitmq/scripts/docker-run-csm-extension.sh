#!/bin/sh

DOCKER_IMAGE="rabbitmq"
DOCKER_IMAGE_TAG="3.6.0-management"
CSM_LOG_LEVEL="debug"
CSM_DEV_MODE="true"
CSM_API_KEY="csm-auth-token"

if [ ! -z ${DOCKER_HOST} ]
then
    export DOCKER_HOST_IP=`echo ${DOCKER_HOST} | cut -d "/" -f 3 | cut -d ":" -f 1`
else
    export DOCKER_HOST_IP=`ip route get 8.8.8.8 | awk 'NR==1 {print $NF}'`
fi

DOCKER_ENDPOINT=http://${DOCKER_HOST_IP}:4445

docker run --name ${CSM_EXTENSION_IMAGE_NAME} \
	-p 8094:8081 \
	-e DOCKER_ENDPOINT=${DOCKER_ENDPOINT} \
	-e DOCKER_IMAGE=${DOCKER_IMAGE} \
	-e DOCKER_IMAGE_TAG=${DOCKER_IMAGE_TAG} \
	-e CSM_LOG_LEVEL=${CSM_LOG_LEVEL} \
	-e CSM_API_KEY=${CSM_API_KEY} \
	-e CSM_DEV_MODE=${CSM_DEV_MODE} \
	-d ${CSM_EXTENSION_IMAGE_NAME}:${CSM_EXTENSION_IMAGE_TAG}