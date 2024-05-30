/* Composite Structs
 */

package composite

import (
	"fmt"
	// "experiment/struct/main" // cannot do this - cyclic dependency
	typeUtil "experiment/type/util"
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
	mp := makeMegaPerson()
	fmt.Println(mp)
	typeUtil.PrintType(mp)

	//typeUtil.PrintType(mp2)
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

type megaPerson person    // not sure what I've created here - I don't think it's a type alias
type megaPerson2 = person // *this* is a type alias

func makeMegaPerson() megaPerson {
	mp := megaPerson{}
	mp.identity.identifier = "1234"
	return mp
}
