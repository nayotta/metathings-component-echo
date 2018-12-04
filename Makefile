DOCKER_BIN=$(shell which docker)

all:
	$(MAKE) -C proto all
	go build --buildmode=plugin -o lib/echo.so pkg/echo/plugin/*.go

protos_from_docker:
	$(DOCKER_BIN) run --rm -v $(PWD):/go/src/github.com/nayotta/metathings-component-echo nayotta/metathings-development /usr/bin/make -C /go/src/github.com/nayotta/metathings-component-echo/proto
