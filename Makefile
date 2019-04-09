DOCKER_BIN=$(shell which docker)

all:
	go build -o metathings-echo cmd/module/main.go

protos_from_docker:
	$(DOCKER_BIN) run --rm -v $(PWD):/go/src/github.com/nayotta/metathings-component-echo nayotta/metathings-development /usr/bin/make -C /go/src/github.com/nayotta/metathings-component-echo/proto
