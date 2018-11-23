package main

import (
	component "github.com/nayotta/metathings/pkg/component"
)

type EchoComponent struct{}

func (self *EchoComponent) Name() string {
	return "echo"
}

func (self *EchoComponent) NewModule(args ...interface{}) (component.Module, error) {
	panic("unimplemented")
}

func NewComponent(args ...interface{}) (component.Component, error) {
	return &EchoComponent{}, nil
}
