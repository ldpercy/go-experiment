/*

Looking for ways to create types dynamically.

A few possibilities...

Type aliases are prefectly substitutable:

	type AliasedType = OriginalType

New named types - not substitutable:

	type NamedType OriginalType

I'm not sure yet whether either of those statement types can be made dynamic - I suspect not.


Interfaces:
https://avilay.rocks/go-dynamic-types/

*/

package dynamic

import "fmt"

type AudioPlayer struct {
	Volume int
}

type Player interface {
	Play(content string)
}

func (ap AudioPlayer) Play(content string) {
	fmt.Printf("Playing audio %v at volume %d\n", content, ap.Volume)
}

var ap = AudioPlayer{Volume: 10}

var p Player = ap
