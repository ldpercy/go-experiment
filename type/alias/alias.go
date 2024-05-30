/*
Type aliasing

So far it appears type aliases are perfectly substitutable.

But new types based on other types aren't compatible, despite having the same structure.
This is so far just fort simple tests on structs, will see how far this travels with other kinds of datatypes.

This might be useful to distinguish (for example) different number types that are fundamentally different kinds of things, like dollars/age/length.

Also appears to not work in reverse, so the new type seems to be truly isolated which is a good thing.

https://go.dev/ref/spec#Type_identity

I think what I'm making here are new 'named types' - which according to the above are always different.

*/

package alias

import (
	"fmt"
)

// Ordinary struct
type Rectangle struct {
	height int
	width  int
}

// Alias of rectangle
type Oblong = Rectangle

// I think this is a new type?
type RightParallelogram Rectangle

func Test() {
	fmt.Println("Type Aliasing")

	var i int = 5
	fmt.Println(i)
	fmt.Println(Squared(i))

	var n number = 12
	fmt.Println(n)
	fmt.Println(Squared(n))

	var d dollars = 10
	fmt.Println(d)
	//fmt.Println(Squared(d)) // also does not work -- good

	//what about in reverse?
	iHave(d)
	//iHave(i) // also does not work - good
}

func TestStructAlias() {
	fmt.Println("Test Struct Aliasing")

	// The ordinary type
	r := Rectangle{height: 4, width: 7}
	fmt.Println(r)
	fmt.Println(Area(r))

	// Now with an alias:
	o := Oblong{height: 11, width: 2}
	fmt.Println(o)
	fmt.Println(Area(o))
	// that works just fine, the function accepts the type alias

	// Now with the new(?) type:
	rp := RightParallelogram{height: 3, width: 8}
	fmt.Println(rp)
	// fmt.Println(Area(rp)) // this does NOT work - the function won't accept this type
}

func Area(r Rectangle) int {
	result := r.height * r.width
	return result
}

/* Let's try a test with simpler types:
 */

type number = int // alias
type dollars int  // new type

func Squared(i int) int {
	result := i * i
	return result
}

// what about in reverse?
// won't accept an int so the typing appears strong here
func iHave(d dollars) {
	fmt.Println("I have dollars:", d)
}
