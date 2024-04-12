package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")
	defer recoverFromPanic()

	var foo = "bar"
	fmt.Println(string(foo))

	dodgyFunction()

	fmt.Println("main end")
}

func dodgyFunction() {
	fmt.Println("dodgyFunction start")
	//defer recoverFromPanic()
	defer fmt.Println("Deferred in dodgyFunction")

	panic("arrrgh!")

	fmt.Println("dodgyFunction end")
}

func recoverFromPanic() {
	r := recover()
	if r != nil {
		fmt.Printf("Panic recovered: %+v \n", r)
	}
	// Make sure to log the error!!
}
