package util

import (
	"fmt"
	"reflect"
)

func Test() {
	fmt.Println("util.Test")
	var foo = "bar"
	fmt.Println(string(foo))
	fmt.Println(TypeOf(foo))
}

// https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go
func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func PrintType(a any) {
	fmt.Println(TypeOf(a))
}
