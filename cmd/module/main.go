package main

import (
	service "github.com/nayotta/metathings-component-echo/pkg/echo/service"
	component "github.com/nayotta/metathings/pkg/component"
)

func main() {
	mdl, err := component.NewModule("echo", new(service.EchoService))
	if err != nil {
		panic(err)
	}

	err = mdl.Launch()
	if err != nil {
		panic(err)
	}
}
