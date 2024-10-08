/*

Looking for ways to create types dynamically.

A few possibilities...

Type aliases are prefectly substitutable:

	type AliasedType = OriginalType

New named types - not substitutable:

	type NamedType OriginalType

I'm not sure yet whether either of those statement types can be made dynamic - I suspect not.


Interfaces (30 Oct 2018):
https://avilay.rocks/go-dynamic-types/

*/

package dynamic

import (
	"fmt"
	"log"
)

func Test() {
	var ap = AudioPlayer{Volume: 10}
	var p Player = ap

	log.Println("p:", p)
	log.Printf("type: %T \n", p) // still reports p as being an AudioPlayer
}

type AudioPlayer struct {
	Volume int
}

type Player interface {
	Play(content string)
}

func (ap AudioPlayer) Play(content string) {
	fmt.Printf("Playing audio %v at volume %d\n", content, ap.Volume)
}
