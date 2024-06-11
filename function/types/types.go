/*
Function types

https://go.dev/ref/spec#Function_types

Experiments with declaring function types, creating instances of types, accepting and returning functions of given type.

https://stackoverflow.com/questions/9398739/working-with-function-types-in-go

	his example is weird, trying a simpler version
*/
package types

import (
	"fmt"
)

func Test() {
	fmt.Println("Function Types")

	fmt.Println(helloEnglish("Function Types"))

	var he Hello = helloEnglish // declare that helloEnglish is of type Hello
	fmt.Println(he("ldpercy"))

	var hs Hello = helloSpanish // declare that helloSpanish is of type Hello
	fmt.Println(hs("ldpercy"))

	// function type as argument type
	fmt.Println(exclamation(he, "ldpercy"))

	// function type as return type
	rh := returnHello()
	fmt.Println(rh("ldpercy"))

	// demonstrate function type identity
	fmt.Println(exclamation(double, "ldpercy"))
}

// Declare the function type
type Hello func(name string) string

// the functions themselves don't declare being of any particular type
func helloEnglish(name string) string {
	return "Hello, " + name
}

func helloSpanish(name string) string {
	return "Hola, " + name
}

// Function type as argument type
func exclamation(h Hello, name string) string {
	return h(name) + "!"
}

// Function type as return type
func returnHello() Hello {
	hf := func(s string) string { return "Hello: " + s }
	return hf
}

// A function that conforms to the signature of Hello but is otherwise unrelated
// This demonstrates function type identity: https://go.dev/ref/spec#Type_identity
func double(s string) string {
	return s + s
}

/* this is the weird bit - a bit of a red herring - I think he's using one function as a receiver for another
func (g Greeting) exclamation(name string) string {
	return g(name) + "!"
}
*/
