BINARY_NAME=bookstore

.PHONY: build

PACKAGE=github.com/guneyin/bookstore
VERSION=$(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH=$(shell git rev-list -1 HEAD)
BUILD_TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S')

LDFLAG_VERSION='${PACKAGE}/common.Version=${VERSION}'
LDFLAG_COMMIT_HASH='${PACKAGE}/common.CommitHash=${COMMIT_HASH}'
LDFLAG_BUILD_TIMESTAMP='${PACKAGE}/common.BuildTime=${BUILD_TIMESTAMP}'

first: tidy vet mock doc run

tidy:
	go mod tidy

vet:
	go vet ./...

mock:
	go run . gen mock

doc:
	swag init

run:
	go run . run

build:
	go build -o ${BINARY_NAME} -ldflags "-X ${LDFLAG_VERSION} -X ${LDFLAG_COMMIT_HASH} -X ${LDFLAG_BUILD_TIMESTAMP}" .

clean:
	go clean
	rm ${BINARY_NAME}


