.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p bin
	go build -o $(shell pwd)/bin/kt-sandbox

.PHONY: build-container
build-container:
	docker build ./docker -t mysql:8.4.6-test

.PHONY: test
test: build-container
	go test -v -count 1 -race ./pkg -ginkgo.v
