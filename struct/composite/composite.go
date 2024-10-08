/* Composite Structs
 */

package composite

import (
	"log"
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
	log.Println("Composite Structs")
	log.Println(makePerson())
	log.Println(makeAnotherPerson())

	pa := personAlias{}
	typeUtil.PrintType(pa)

	mp := makeMegaPerson()
	log.Println(mp)
	typeUtil.PrintType(mp)
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

// Alisaing

type personAlias = person // *this* is a type alias

/*
not sure what I've created here - I don't think it's a type alias
It seems to be brand new type, that just happens to be duplicate of person
Looks to be a new 'named type' see experiment/type/alias for more info.
*/
type megaPerson person

func makeMegaPerson() megaPerson {
	mp := megaPerson{}
	mp.identity.identifier = "1234"
	return mp
}
