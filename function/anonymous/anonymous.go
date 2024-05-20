/*
some simple experiments with anonymous functions
*/

package anonymous

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
