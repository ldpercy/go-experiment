package main

import (
	"fmt"
	"reflect"
)

func modifyP1(p person) {
	p.name = "Bill"
}

func modifyP2(p *person) {
	p.name = "Graham"
}

func modifyString1(s string) string {
	return fmt.Sprintf("%v %v", s, s)
}

//util

func print(name string, variable any) {
	fmt.Printf("%v: %#v \n", name, variable)
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
