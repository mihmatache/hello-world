SHELL:=/bin/bash
APP:=hello-world
TOP_DIR:=$(notdir $(CURDIR))
BIN_DIR:=_bin
PORT?=5000
DOCKER_REPO?="mmatache"
IMAGE?=$(DOCKER_REPO)/$(APP)
SERVER_TYPE?=grpc
ADDRESS?="127.0.0.1"
TIMEOUT?=5
CHART_NAME?=$(APP)


ifeq ($(VERSION),)
	VERSION:=$(shell git describe --tags --dirty --always)
endif


all: install-go-tools lint build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP)

install:
	go install
	
install-go-tools:
	@./scripts/install_tools.sh
	go install github.com/golang/mock/mockgen

lint:
	golangci-lint run ./...

generate:
	go generate -v ./...

run-server: build
	$(BIN_DIR)/$(APP) grpc server --port $(PORT)

run-client: build
	$(BIN_DIR)/$(APP) grpc client --timeout $(TIMEOUT) 127.0.0.1:5000

image: build
	docker build -t $(IMAGE):$(VERSION) .

push-images: image
	docker push $(IMAGE):$(VERSION) 

helm-install:
	helm install --set image.tag=$(VERSION) --replace $(CHART_NAME) --set service.name=$(CHART_NAME) --set service.Port=$(PORT) --set serverType=$(SERVER_TYPE) ./deployments/helm/$(APP)

helm-uninstall:
	helm uninstall $(APP)
