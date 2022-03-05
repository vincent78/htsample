GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOBIN = ./build/bin

NAME=htsample
BINARY_NAME=${NAME}
BINARY_UNIX=$(BINARY_NAME)_unix

#IMAGE_REGISTRY=registry.cn-hangzhou.aliyuncs.com/quicktron_robot
IMAGE_REGISTRY=127.0.0.1/htsample
IMAGE_NAME=$(IMAGE_REGISTRY)/$(NAME)

GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
BUILD_DATE=$(shell date "+%y-%m-%dT%H:%M:%S")

CGO_ENABLED=0
LDFLAGS=-X global.GitCommit=$(GIT_TAG) -X global.BuildDate=$(BUILD_DATE)


BUILD_TAG=$(GIT_COMMIT)
#IMAGE_TAG=$(shell date "+%y%m%d%H%M")
IMAGE_TAG=dev_$(shell date "+%m%d")
IMAGE_TAG_RELEASE=1.0.0


all: clean runServer

build:
	go mod tidy
	$(GOBUILD) -a -ldflags "-w -s ${LDFLAGS}" -o $(GOBIN)/$(BINARY_NAME) -v

runServer:
	$(GOBUILD) -a -ldflags "-w -s ${LDFLAGS}" -o $(GOBIN)/$(BINARY_NAME) -v
	$(GOBIN)/${BINARY_NAME} --port 8080 server


linux:
	go mod tidy
	$(GOBUILD) -a -ldflags "-w -s ${LDFLAGS}" -o $(GOBIN)/$(BINARY_NAME) -v
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GOBUILD) -a -ldflags "-w -s ${LDFLAGS}" -o $(GOBIN)/$(BINARY_NAME) -v


docker: build
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG)  .
	#docker push $(IMAGE_NAME):$(IMAGE_TAG)


dockerRelease: linux
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG_RELEASE)  .
#	docker push $(IMAGE_NAME):$(IMAGE_TAG_RELEASE)


test:
	go test -v ./...

clean:
	rm -rf ./$(NAME) && rm -rf ./build


.PHONY: build