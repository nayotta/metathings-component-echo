package main

import "reflect"

import cmp "github.com/nayotta/metathings-component-echo/pkg/echo/service"

func main() {
	x := new(cmp.EchoModule)
	rx := reflect.ValueOf(x)
	fn := rx.MethodByName("Echo")
	println(fn.Type().NumIn())
}
