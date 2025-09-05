KUBERNETES_VERSION = 1.33.1
CERT_MANAGER_VERSION = 1.15.3
MOCO_VERSION = 0.28.0

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p bin
	go build -o $(shell pwd)/bin/kt-moco-chaos

.PHONY: start
start: install-aqua
	kind create cluster --name=sandbox --config=./kind/kind-config.yaml --image=kindest/node:v$(KUBERNETES_VERSION) --wait 1m
	kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v$(CERT_MANAGER_VERSION)/cert-manager.yaml
	kubectl -n cert-manager wait --for=condition=available --timeout=180s --all deployments
	kubectl apply -f https://github.com/cybozu-go/moco/releases/download/v$(MOCO_VERSION)/moco.yaml
	kubectl -n moco-system wait --for=condition=available --timeout=180s --all deployments

.PHONY: stop
stop: install-aqua
	kind delete cluster --name sandbox

.PHONY: setup-aqua
setup-aqua:
	go install github.com/aquaproj/aqua/v2/cmd/aqua@latest

.PHONY: install-aqua
install-aqua:
	aqua i
