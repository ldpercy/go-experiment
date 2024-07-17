/* function
 */
package function

import "log"

func Test() {

	name := "Alice"
	log.Println(name)
	log.Println(double(name))
	log.Println(name)

	f := "foobar"
	acceptsReference(&f)
	log.Println(f)
}

func double(s string) string {
	s = s + s
	return s
}

// acceptsReference
/*
At the moment this weirds me out a bit - to say that a function accepts a reference you write the argument as

	argname *type

whereas it feels like it should be

	argname &type

Then it would make sense when dereferencing it in the function body
*/
func acceptsReference(r *string) {

	log.Println(r)
	*r = "asdf"
}

// func test(thing &string) { ... }
// And this sort of argument syntax doesn't seem to be used..?

// Returning references

func returnReference() *string {
	s := "qwerty"
	return &s
}

// again, this looks the wrong way around, the return type should &string

// func returnReference2() &string { ...}
// once more this kind of syntax appears unused

/* I wonder if there's some other reason why they didn't use the ampersand in the function signatures?
A parsing thing or a C thing?
Or maybe there's something I'm not getting yet.

This summarizes it fairly well:
https://stackoverflow.com/questions/3552626/what-does-the-asterisk-do-in-go

But it still should have been an ampersand for types.


*/
