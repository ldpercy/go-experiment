/*
https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go
https://golangtutorial.dev/tips/find-type-of-an-object/
*/

package check

import (
	"fmt"
	"log"
	"reflect"
)

func Test() {
	log.Println("Simple string:")
	var foo string
	foo = "bar"
	log.Println(SprintType(foo))                   // string
	log.Println(ReflectTypeOf(foo))                // string
	log.Println(ReflectTypeOf(ReflectTypeOf(foo))) // *reflect.rtype - not yet sure what this is
	log.Println(ReflectKind(foo))                  // string

	log.Println("Type alias:")
	type number = int
	var five number
	five = 5
	log.Println(SprintType(five))    // int
	log.Println(ReflectTypeOf(five)) // int
	log.Println(ReflectKind(five))   // int
	// aliases return their source types

	log.Println("Named Type:")
	type dollars int
	var lobster dollars
	lobster = 20
	log.Println(SprintType(lobster))    // check.dollars
	log.Println(ReflectTypeOf(lobster)) // check.dollars
	log.Println(ReflectKind(lobster))   // int
	// named types are returned except for kind
}

func SprintType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func ReflectTypeOf(variable interface{}) any {
	return reflect.TypeOf(variable)
}

func ReflectKind(variable interface{}) any {
	return reflect.ValueOf(variable).Kind()
}
