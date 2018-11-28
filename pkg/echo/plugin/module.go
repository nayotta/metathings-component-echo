package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"

	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
)

type EchoModuleOption struct {
}

type EchoModule struct {
	logger log.FieldLogger
	opt    *EchoModuleOption
	srv    *service.EchoService
	app    *fx.App
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

func SetupEchoModule(mdl *EchoModule, logger log.FieldLogger, opt *EchoModuleOption, srv *service.EchoService) {
	mdl.logger = logger
	mdl.opt = opt
	mdl.srv = srv
}
