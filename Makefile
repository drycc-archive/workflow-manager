SHORT_NAME := workflow-manager

include versioning.mk

DEV_ENV_IMAGE := quay.io/drycc/go-dev:v0.22.0
SWAGGER_IMAGE := quay.io/goswagger/swagger:v0.18.0
DEV_ENV_WORK_DIR := /go/src/github.com/drycc/${SHORT_NAME}
DEV_ENV_CMD := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR} ${DEV_ENV_IMAGE}
SWAGGER_CMD := docker run --rm -e GOPATH=/go -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR} ${SWAGGER_IMAGE}
SHELL_SCRIPTS = rootfs/bin/doctor

# Common flags passed into Go's linker.
LDFLAGS := "-s -X main.version=${VERSION}"

# Docker Root FS
BINDIR := ${CURDIR}/rootfs/bin

# Legacy support for DEV_REGISTRY, plus new support for DRYCC_REGISTRY.
ifdef ${DEV_REGISTRY}
  DRYCC_REGISTRY = ${DEV_REGISTRY}
endif

all: build docker-build docker-push

# Containerized dependency resolution / initial workspace setup
bootstrap:
	${DEV_ENV_CMD} glide install

# This illustrates a two-stage Docker build. docker-compile runs inside of
# the Docker environment. Other alternatives are cross-compiling, doing
# the build as a `docker build`.
build:
	mkdir -p ${BINDIR}
	${DEV_ENV_CMD} go build -o rootfs/bin/boot -ldflags ${LDFLAGS} boot.go

swagger-clientstub:
	${SWAGGER_CMD} generate client -A WorkflowManager -t pkg/swagger -f api/swagger-spec/swagger.yml

test:
	${SWAGGER_CMD} validate ./api/swagger-spec/swagger.yml
	${DEV_ENV_CMD} sh -c 'go test -v $$(glide nv)'

test-cover:
	${DEV_ENV_CMD} test-cover.sh

test-style:
	${DEV_ENV_CMD} shellcheck $(SHELL_SCRIPTS)

# For cases where we're building from local
# We also alter the RC file to set the image name.
docker-build: build
	docker build ${DOCKER_BUILD_FLAGS} -t ${IMAGE} rootfs
	docker tag ${IMAGE} ${MUTABLE_IMAGE}
