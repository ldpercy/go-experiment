/*
	Pointers

This is based on:
https://stackoverflow.com/questions/9993178/is-there-a-nice-way-to-simulate-a-maybe-or-option-type-in-go

Want to see what the difference is between having an 'ordinary' substruct, and one done with pointers

Ref: https://stackoverflow.com/questions/59964619/difference-using-pointer-in-struct-fields

Off the top of my head I guess using pointers might be a little bit more lightweight in terms of memory (but probably not by a lot).

Maybe not:
https://stackoverflow.com/questions/59964619/difference-using-pointer-in-struct-fields
*/
package composite

import "log"

func TestPointers() {
	log.Println("Test Pointers")
	po := parentOrdinary{
		name:     "po",
		children: []child{{name: "poc1"}, {name: "poc2"}},
	}

	log.Printf("%#v \n", po)

	ac := []child{{name: "c1"}, {name: "c2"}}

	pp := parentPointer{
		name:     "p2",
		children: &ac,
	}
	//fmt.Printf("%#v \n", ac)
	log.Printf("%#v \n", pp)
}

type parentOrdinary struct {
	name     string
	children []child
}

type parentPointer struct {
	name     string
	children *[]child
}

type child struct {
	name string
}
