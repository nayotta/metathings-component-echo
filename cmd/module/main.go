package main

import (
	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
	component "github.com/nayotta/metathings/pkg/component"
)

var Version string

func main() {
	mdl, err := component.NewModule(
		component.SetTarget(new(service.EchoService)),
		component.SetVersion(Version),
	)
	if err != nil {
		panic(err)
	}

	err = mdl.Launch()
	if err != nil {
		panic(err)
	}
}
