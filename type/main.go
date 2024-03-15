package main

import (
	"fmt"
	"reflect"
)

func main() {

	var foo = "bar"

	fmt.Println(string(foo))

	fmt.Println(typeof(foo))

}

// https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
