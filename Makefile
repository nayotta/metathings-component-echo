all:
	protoc -I. -I/usr/local/include -I$(GOPATH)/src --go_out=plugins=grpc:. proto/service.proto
	go build --buildmode=plugin -o lib/echo.so pkg/component.go
