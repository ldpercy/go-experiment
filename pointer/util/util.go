package util

import (
	"fmt"
	"reflect"
)

//util

func Print(name string, variable any) {
	fmt.Printf("%v: %#v \n", name, variable)
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
