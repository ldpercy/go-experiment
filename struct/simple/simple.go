package simple

import (
	"log"
)

type simpleStruct struct {
	name1 string
	name2 string
}

func Test() {
	log.Println("Simple Structs")

	log.Println(makeStruct())
}

func makeStruct() simpleStruct {
	myStruct := simpleStruct{
		name1: "foo",
		name2: "bar",
	}
	return myStruct
}

// Empty struct??
type emptyStruct struct{}

func testEmptyStruct() emptyStruct {
	es := emptyStruct{}
	return es
}
