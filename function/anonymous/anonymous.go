/*
some simple experiments with anonymous functions
*/

package anonymous

import (
	"fmt"
)

func Test() {
	fmt.Println("Anonymous Functions")

	fmt.Println(foo())
	fmt.Println(bar())
	rf := retFunc()
	fmt.Println(rf())
}

var foo = func() string {
	return "foo"
}

var bar = func() string {
	return "bar"
}

var retFunc = func() func() string {
	result := func() string { return "blah" }
	return result
}
