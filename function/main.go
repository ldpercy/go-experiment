package main

import (
	_ "experiment/function/anonymous"
	"experiment/function/multireturn"
	"fmt"
)

func main() {

	/*
		// Anonymous functions

		fmt.Println(anon.Foo())
		fmt.Println(anon.Bar())
		rf := anon.RetFunc()
		fmt.Println(rf())
	*/
	/*
		doesn't look like you can loop over an import

		for index, value := range anon {
			fmt.Printf("%d is index, %d is value", index, value)
		}
	*/

	// Multiple Return
	fmt.Println(multireturn.ReturnOne(5))
	//fmt.Println(multireturn.ReturnOneBracketed(6))
	fmt.Println(multireturn.ReturnTwo(7))
	fmt.Println(multireturn.ReturnThree(8))

	a, b := multireturn.ReturnTwoFuncs()
	fmt.Println(a(11), b(12))
}
