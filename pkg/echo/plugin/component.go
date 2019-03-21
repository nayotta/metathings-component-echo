package main

import (
	"context"
	"strings"
	"time"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	component "github.com/nayotta/metathings/pkg/component"
	component_pb "github.com/nayotta/metathings/pkg/proto/component"
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"go.uber.org/fx"

	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
	pb "github.com/nayotta/metathings-component-echo/proto"
)

type RunModuleOption struct {
	cmd_contrib.ModuleBaseOption `mapstructure:",squash"`
	Heartbeat                    cmd_contrib.HeartbeatOption `mapstruct:"heartbeat"`
}

func CreateRunModuleOption() RunModuleOption {
	return RunModuleOption{
		ModuleBaseOption: cmd_contrib.CreateModuleBaseOption(),
	}
}

func NewEchoModuleOption(component component.Component, opt *RunModuleOption) *EchoModuleOption {
	return &EchoModuleOption{
		Name:              opt.GetName(),
		Component:         component.Name(),
		HeartbeatInterval: time.Duration(opt.Heartbeat.GetInterval()) * time.Second,
	}
}

func NewEchoServiceOption(opt *RunModuleOption) *service.EchoServiceOption {
	return &service.EchoServiceOption{}
}

type EchoComponent struct{}

func (self *EchoComponent) Name() string {
	return "echo"
}

func (self *EchoComponent) NewModule(args []string) (component.Module, error) {
	mdl := &EchoModule{}

	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			func() component.Component { return self },
			func() (*RunModuleOption, error) {
				tmp_opt := CreateRunModuleOption()
				var opt *RunModuleOption = &tmp_opt
				var err error

				flags := pflag.NewFlagSet("echo", pflag.ExitOnError)
				flags.StringVarP(opt.GetConfigP(), "config", "c", "", "Config file")
				flags.StringVar(opt.GetNameP(), "name", "echo", "Module name")
				flags.StringVarP(opt.GetListenP(), "listen", "l", "127.0.0.1:13401", "MetaThings Echo Service listening address")
				flags.IntVar(opt.Heartbeat.GetIntervalP(), "heartbeat-interval", 15, "MetaThings Echo Service heartbeat interval (seconds)")
				if err = flags.Parse(args); err != nil {
					return nil, err
				}

				if opt.GetConfig() != "" {
					v := viper.New()
					v.AutomaticEnv()
					v.SetEnvPrefix(component.METATHINGS_COMPONENT_PREFIX)
					v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
					v.BindEnv("stage")

					v.SetConfigFile(opt.GetConfig())
					if err = v.ReadInConfig(); err != nil {
						return nil, err
					}

					tmp1_opt := CreateRunModuleOption()
					cmd_helper.UnmarshalConfig(&tmp1_opt, v)
					opt = &tmp1_opt
				}

				return opt, nil
			},
			func(opt *RunModuleOption) (
				cmd_contrib.LoggerOptioner,
				cmd_contrib.ListenOptioner,
				cmd_contrib.TransportCredentialOptioner,
				cmd_contrib.ServiceEndpointsOptioner,
			) {
				return opt, opt, opt, opt
			},
			cmd_contrib.NewLogger("echo"),
			NewEchoServiceOption,
			NewEchoModuleOption,
			service.NewEchoService,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			func(logger log.FieldLogger, opt *EchoModuleOption, srv *service.EchoService, cli_fty *client_helper.ClientFactory, lc fx.Lifecycle) *EchoModule {
				mdl.logger = logger
				mdl.opt = opt
				mdl.srv = srv
				mdl.cli_fty = cli_fty

				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						go mdl.heartbeat_loop()
						return nil
					},
				})

				return mdl
			},
			func(m *EchoModule, logger log.FieldLogger) (pb.EchoServiceServer, component_pb.ModuleServiceServer) {
				return m.srv, component.NewGrpcModuleWrapper(m.srv, logger)
			},
		),
		fx.Invoke(
			pb.RegisterEchoServiceServer,
			component_pb.RegisterModuleServiceServer,
		),
	)
	mdl.app = app

	return mdl, nil
}

var NewComponent component.NewComponent = func() (component.Component, error) {
	return &EchoComponent{}, nil
}
