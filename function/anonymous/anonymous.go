/*
some simple experiments with anonymous functions
*/

package anonymous

import (
	"fmt"
)

func Test() {

	// Anonymous functions

	fmt.Println(Foo())
	fmt.Println(Bar())
	rf := RetFunc()
	fmt.Println(rf())
}

var Foo = func() string {
	return "foo"
}

var Bar = func() string {
	return "bar"
}

var RetFunc = func() func() string {

	result := func() string { return "blah" }
	return result
}
