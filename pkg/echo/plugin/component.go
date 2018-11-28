package main

import (
	"fmt"
	"strings"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	component "github.com/nayotta/metathings/pkg/component"
	component_pb "github.com/nayotta/metathings/pkg/proto/component"
	"github.com/nayotta/viper"
	"github.com/spf13/pflag"
	"go.uber.org/fx"

	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
	pb "github.com/nayotta/metathings-component-echo/proto"
)

type RunModuleOption struct {
	cmd_contrib.ModuleBaseOption `mapstructure:",squash"`
}

func CreateRunModuleOption() RunModuleOption {
	return RunModuleOption{
		ModuleBaseOption: cmd_contrib.CreateModuleBaseOption(),
	}
}

func NewEchoModuleOption(opt *RunModuleOption) *EchoModuleOption {
	return &EchoModuleOption{}
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
		fx.Provide(
			func() *EchoModule { return mdl },
			func() (*RunModuleOption, error) {
				tmp_opt := CreateRunModuleOption()
				var opt *RunModuleOption = &tmp_opt
				var err error

				flags := pflag.NewFlagSet("echo", pflag.ExitOnError)
				flags.StringVarP(opt.GetConfigP(), "config", "c", "", "Config file")
				flags.StringVar(opt.GetNameP(), "name", "echo", "Module name")
				flags.StringVarP(opt.GetListenP(), "listen", "l", "127.0.0.1:13401", "MetaThings Echo Service listening address")
				if err = flags.Parse(args); err != nil {
					return nil, err
				}

				fmt.Printf("!!!! %v\n", cmd_helper.GetStageFromEnv())

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
			) {
				return opt, opt, opt
			},
			cmd_contrib.NewLogger("echo"),
			NewEchoServiceOption,
			NewEchoModuleOption,
			service.NewEchoService,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			func(m *EchoModule) (pb.EchoServiceServer, component_pb.ModuleServiceServer) {
				return m.srv, component.NewGrpcModuleWrapper(m.srv)
			},
		),
		fx.Invoke(
			pb.RegisterEchoServiceServer,
			component_pb.RegisterModuleServiceServer,
			SetupEchoModule,
		),
	)
	mdl.app = app

	return mdl, nil
}

var NewComponent component.NewComponent = func() (component.Component, error) {
	return &EchoComponent{}, nil
}
