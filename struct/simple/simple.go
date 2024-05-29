package simple

import (
	"fmt"
)

type simpleStruct struct {
	name1 string
	name2 string
}

func Test() {
	fmt.Println("Simple Structs")

	fmt.Println(makeStruct())
}

func makeStruct() simpleStruct {
	myStruct := simpleStruct{
		name1: "foo",
		name2: "bar",
	}
	return myStruct
}
