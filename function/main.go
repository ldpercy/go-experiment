package main

import (
	_ "experiment/function/anonymous"
	_ "experiment/function/multireturn"
	"experiment/function/typing"
)

func main() {

	/*
		doesn't look like you can loop over an import

		for index, value := range anon {
			fmt.Printf("%d is index, %d is value", index, value)
		}
	*/
	// anonymous.Test()
	// multireturn.Test()
	typing.Test()
}
