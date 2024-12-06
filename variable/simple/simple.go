package simple

import "log"

/*
var a     // cannot do this, must be typed or have a value, syntax error
*/

// uninitialised variable
var b int // gets default value, explicitly typed

// initialised variables
var name1 = "Alice"      // implictly typed
var name2 string = "Bob" // explicitly typed

func Test() {
	log.Println("Simple Test")

	log.Println(b)
	log.Println(name1)
	log.Println(name2)

	name3 := "Carol" // shorthand is implictly typed
	log.Println(name3)

	//	name4 string := "Dave"		// cannot use a type with a shorthand declaration

}
