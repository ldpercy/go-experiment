package main

import (
	anon "experiment/function/anonymous"
	"fmt"
)

func main() {

	fmt.Println(anon.Foo())
	fmt.Println(anon.Bar())

	rf := anon.RetFunc()
	fmt.Println(rf())

	/*
		doesn't look like you can loop over an import

		for index, value := range anon {
			fmt.Printf("%d is index, %d is value", index, value)
		}
	*/

}
