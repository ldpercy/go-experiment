package main

import "fmt"

var mainAsWell = "mainAsWell hello"

// fmt.Println("funcInFolder") // not allowed, no non-decs outside funcs

func funcAsWell() {
	fmt.Println(mainAsWell)
	//mainVar := "local version of mainVar"
	fmt.Println(mainVar)
	mainVar = "Will this overwrite the local or the package version?"
	/*
		overwrites local if there is one
		overwrites package var otherwise
	*/
	fmt.Println(mainVar)
	var packageVar = "this might be a package var" // but it's not, it's local
	fmt.Println(packageVar)
}

/*
func testPackageVar() {
	fmt.Println(packageVar)
}
*/
