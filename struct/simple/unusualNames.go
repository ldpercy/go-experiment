package simple

func Test2() {

}

/*
Was to see if Go can do funky var names by using quotes - this works in some langauges.
I don't think this will be possible in Go.
I'm pretty sure all variables must start with a letter so they can be capitalised for public.

Apparently you can start with an underscore too, though I imagine these will be private by default.

*/

type oddStruct struct {
	_foo string
	bar  string
}
