/*
	Pointers

This is based on:
https://stackoverflow.com/questions/9993178/is-there-a-nice-way-to-simulate-a-maybe-or-option-type-in-go

Wnat to see what the difference is between having an 'ordinary' substruct, and one done with pointers
*/
package composite

import "fmt"

func TestPointers() {

	p1 := parent{
		name: "p1",
	}

	fmt.Printf("%#v \n", p1)

	ac := []child{{name: "c1"}, {name: "c2"}}

	p2 := parent{
		name:     "p2",
		children: &ac,
	}
	fmt.Printf("%#v \n", ac)
	fmt.Printf("%#v \n", p2)
}

type parent struct {
	name     string
	children *[]child
}

type child struct {
	name string
}
