DOCKER_BIN=$(shell which docker)

all:
	go build -ldflags "-X=google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o metathings-echo cmd/module/main.go

protos:
	$(DOCKER_BIN) run --rm -v $(PWD):/go/src/github.com/nayotta/metathings-component-echo --entrypoint /go/src/github.com/nayotta/metathings-component-echo/scripts/build-proto.sh jaegertracing/protobuf
