#! /bin/bash
mkdir -p /go/src/github.com/nayotta/metathings-component-echo/pbset

protoc \
	-I /go/src \
	-I /go/src/github.com/nayotta/metathings-component-echo \
	-I /go/src/github.com/nayotta/metathings-component-echo/vendor \
	--include_imports \
	--descriptor_set_out=/go/src/github.com/nayotta/metathings-component-echo/pbset/service.pbset \
	--go_out=plugins=grpc:/go/src/ \
	/go/src/github.com/nayotta/metathings-component-echo/proto/*.proto
