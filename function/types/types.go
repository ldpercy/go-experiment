/*
Function types

https://go.dev/ref/spec#Function_types

Experiments with declaring function types, creating instances of types, accepting and returning functions of given type.
*/
package types

import (
	"fmt"
)

func Test() {
	fmt.Println(english("ANisus"))
	fmt.Println(english.exclamation("ANisus"))
}

type Greeting func(name string) string

func (g Greeting) exclamation(name string) string {
	return g(name) + "!"
}

var english = Greeting(func(name string) string {
	return "Hello, " + name
})
