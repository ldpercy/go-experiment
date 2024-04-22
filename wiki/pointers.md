Pointers
========


Need to get passing around pointers/references sorted out.


Pass by value types
-------------------

https://www.youtube.com/watch?v=LBLN4Wu5U4w&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=13

Group A: Non-pointer values
	string
	int
	float
	boolean
	array
	struct

	These are pass-by-value

	Pointers themselves are probably in this group

Group B: Pointer-wrapper values
	slice
	map
	function

	These effectively get passed by reference


Pointers
--------

Pointer to variable:

	valueName := "foobar"

	myPointer := &valueName 			// is pointer to valueName

	// 0x0123456789ab

Dereference pointer to get underlying value at its memory address:

	underlyingValue := *myPointer

Function that accepts a pointer:

	func updateValue(p *valueType) {
		*p := someNewValue
	}

(This explicit pointer syntax is implicit for pointer-wrapper types.)



passing structs into functions
------------------------------


```golang
	func editStruct(foo MyStruct)
	{
		// foo is passed in as a copy
		foo.thing = "newvalue"
		// foo will be thrown away here and the edit lost
	}


	func editStruct(foo *MyStruct)
	{
		// foo is passed in as a pointer to the original
		foo.thing = "newvalue"
		// edits are made to the original and will persist
	}


	// Or, in function receiver syntax:

	func (foo *MyStruct) editStruct()
	{
		// foo is passed in as a pointer to the original
		foo.thing = "newvalue"
		// edits are made to the original and will persist
	}

	// The latter two must be called like this:


	editStruct(&myFoo)


```