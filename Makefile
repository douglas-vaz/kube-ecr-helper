.PHONY: all
all: build-deps build fmt
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/kube-ecr"
setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
build-deps:
	dep ensure
update-deps:
	dep ensure
build:  build-deps compile fmt
compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)
install:
	go install ./...
fmt:
	go fmt $(go list | grep -v vendor)
