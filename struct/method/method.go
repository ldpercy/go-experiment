/* Struct Methods
aka
receivers

https://gobyexample.com/methods
https://medium.com/@senowijayanto/understanding-golang-receivers-when-to-use-pointers-and-value-receivers-58f9e346ed65


See also
https://stackoverflow.com/questions/59964619/difference-using-pointer-in-struct-fields
For notes about idiomatic use of pointers in receivers

*/

package method

import "log"

func Test() {

	log.Println("Struct methods")

	r := rect{width: 10, height: 5}

	log.Println("area: ", r.area())
	log.Println("perim:", r.perim())

	rp := &r
	log.Println("area: ", rp.area())
	log.Println("perim:", rp.perim())
}

type rect struct {
	width, height int
}

// pointer receiver
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/* Okay need to figure these out - from what I've read so far the pointer receiver is the standard way of doing things
What is the value receiver used for?
Probably comes down to pass by value vs pass by reference.
I'm guessing the value receiver method will receive a copy of the struct, so mutations etc won't affect the original.
Whereas the pointer receiver will presumably get the original.

So if you need mutators I think you'd want pointer receivers.
For constructors, and maybe simple getters, value receivers should be okay.
Will try and revise as I go.

See also:
https://go.dev/wiki/CodeReviewComments#receiver-type

*/
