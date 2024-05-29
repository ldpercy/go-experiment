/* Composite Structs
 */

package composite

import (
	"fmt"
)

type identity struct {
	name       string
	identifier string
}

type address struct {
	number uint
	street string
	suburb string
}

type person struct {
	identity identity
	address  address
}

/* Test */
func Test() {
	fmt.Println("Composite Structs")
	fmt.Println(makePerson())
	fmt.Println(makeAnotherPerson())
	fmt.Println(makeMegaPerson())
}

func makePerson() person {
	thisPerson := person{
		identity: identity{name: "Phil"},
	}
	return thisPerson
}

func makeAnotherPerson() person {
	thisPerson := person{}
	thisPerson.identity.name = "Tony"
	return thisPerson
}

type megaPerson person // not sure what I've created here - I don't think it's a type alias

func makeMegaPerson() megaPerson {
	mp := megaPerson{}
	mp.identity.identifier = "1234"
	return mp
}
