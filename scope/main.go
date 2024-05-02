/*
Scope
=====

Figure out how scope works in go across:
* files
* packages
* functions
* lambdas
*/
package main

import (
	"experiment/scope/folder"
	"fmt"
)

var mainVar = "this is main"

func main() {
	fmt.Println("main")
	fmt.Println(mainVar)
	fmt.Println(mainAsWell)
	funcAsWell()
	fmt.Println(mainVar)
	// fmt.Println(packageVar)

	folder.FolderStuff()

}
