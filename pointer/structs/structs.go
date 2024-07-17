package structs

import (
	"experiment/pointer/util"
	"fmt"
)

type person struct {
	name string
	age  int
}

type classroom []person

var p1 = person{
	name: "Larry",
	age:  33,
}

var c1 = classroom{
	person{
		name: "Larry",
		age:  33,
	},
	person{
		name: "Curly",
		age:  44,
	},
	person{
		name: "Moe",
		age:  55,
	},
}

func Test() {

	util.Print("p1", p1)

	modifyP1(p1)
	util.Print("p1", p1)
	/*
		modifyP2(&p1)
		print("p1", p1)

		fmt.Printf("&p1: %#v \n", &p1)

		ptr1 := &p1.name
		fmt.Printf("ptr1: %#v \n", ptr1)

		print("typeof(ptr1)", typeof(ptr1))
	*/

	//newString := "foo"
	//newString1 := modifyString1(newString)
	//fmt.Printf("newString1: %#v \n", newString1)

	// can you point to submembers?
	util.Print("c1", c1)
	util.Print("c1[0]", c1[0])

	modifyP2(&c1[0])
	util.Print("c1", c1)
	// seems so

}

func modifyP1(p person) {
	p.name = "Bill"
}

func modifyP2(p *person) {
	p.name = "Graham"
}

func modifyString1(s string) string {
	return fmt.Sprintf("%v %v", s, s)
}
