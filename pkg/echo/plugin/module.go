package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	device_pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"

	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
)

type EchoModuleOption struct {
	Name              string
	Component         string
	HeartbeatInterval time.Duration
}

type EchoModule struct {
	logger  log.FieldLogger
	opt     *EchoModuleOption
	srv     *service.EchoService
	app     *fx.App
	cli_fty *client_helper.ClientFactory
}

func (self *EchoModule) Start() error {
	return self.app.Start(context.Background())
}

func (self *EchoModule) Stop() error {
	return self.app.Stop(context.Background())
}

func (self *EchoModule) Err() error {
	return self.app.Err()
}

func (self *EchoModule) heartbeat_loop() {
	for {
		go self.heartbeat_once()
		time.Sleep(self.opt.HeartbeatInterval)
	}
}

func (self *EchoModule) heartbeat_once() {
	cli, cfn, err := self.cli_fty.NewDeviceServiceClient()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to connect to device")
		return
	}
	defer cfn()

	req := &device_pb.HeartbeatRequest{
		Module: &deviced_pb.OpModule{
			Name:      &wrappers.StringValue{Value: self.opt.Name},
			Component: &wrappers.StringValue{Value: self.opt.Component},
		},
	}
	_, err = cli.Heartbeat(context.Background(), req)
	if err != nil {
		self.logger.WithError(err).Warningf("failed to heartbeat to device")
		return
	}

	self.logger.WithFields(log.Fields{}).Debugf("heartbeat")
}
