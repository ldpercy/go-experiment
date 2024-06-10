/*
Function types

https://go.dev/ref/spec#Function_types

Experiments with declaring function types, creating instances of types, accepting and returning functions of given type.

https://stackoverflow.com/questions/9398739/working-with-function-types-in-go

	his exmaple is weird, trying a simpler version
*/
package types

import (
	"fmt"
)

func Test() {
	fmt.Println("Function Types")
	fmt.Println(english("ldpercy"))
	//fmt.Println(english.exclamation("ANisus"))
}

type Greeting func(name string) string // this declares the function type

var english = Greeting(func(name string) string {
	return "Hello, " + name
})

func main() {
	var g Greeting
	g = english
	fmt.Println(g("ANisus"))
}

/* this is the weird bit - a bit of a red herring - I think he's using one function as a receiver for another
func (g Greeting) exclamation(name string) string {
	return g(name) + "!"
}
*/
