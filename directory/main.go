package main

import (
	"fmt"
	"io/fs"
)

func main() {

	directory := "."

	// readDir
	//fmt.Println("osReadDir")
	//printFiles(osReadDir(directory))

	//fmt.Println("fsReadDir")
	//printFiles(fsReadDir(directory))

	// walkDir
	fmt.Println("walkDir")
	getFiles(directory)
	//result := getFiles(directory)
	//printFiles(result)

}

/*
https://pkg.go.dev/io/fs#ReadDir
https://pkg.go.dev/io/ioutil#ReadDir		deprecated!!
https://pkg.go.dev/os#ReadDir

*/

func printFiles(files []fs.DirEntry) {
	for _, file := range files {
		fmt.Printf("%#v \n", file)
	}
}
