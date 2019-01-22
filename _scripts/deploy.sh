#!/usr/bin/env bash
#
# Build and push Docker images to Docker Hub and quay.io.
#

cd "$(dirname "$0")" || exit 1
docker login -e="$DOCKER_EMAIL" -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD"
DRYCC_REGISTRY='' make -C .. docker-push

# in order to not build the container again, let's do some
# docker tagging trickery.
version="git-$(git rev-parse --short HEAD)"
docker tag drycc/workflow-manager:canary quay.io/drycc/workflow-manager:canary
docker tag drycc/workflow-manager:${version} quay.io/drycc/workflow-manager:${version}

docker login -e="$QUAY_EMAIL" -u="$QUAY_USERNAME" -p="$QUAY_PASSWORD" quay.io
DRYCC_REGISTRY=quay.io/ make -C .. docker-push
