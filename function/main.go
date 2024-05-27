package main

import (
	_ "experiment/function/anonymous"
	_ "experiment/function/multireturn"
	"experiment/function/types"
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
	types.Test()
}
