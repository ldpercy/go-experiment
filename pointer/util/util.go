package util

import (
	"log"
	"reflect"
)

//util

func Print(name string, variable any) {
	log.Printf("%v: %#v \n", name, variable)
}

func Typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
