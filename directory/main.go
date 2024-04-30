package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {

	directory := ".."

	fmt.Println("osReadDir")
	printFiles(osReadDir(directory))

	fmt.Println("fsReadDir")
	printFiles(fsReadDir(directory))

	/*
		// read the file
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}

		// output some info
		fmt.Println(reflect.TypeOf(content)) // content is a slice
		fmt.Println(file.Stat())

		fmt.Println(string(content))
	*/
}

/*
https://pkg.go.dev/io/fs#ReadDir
https://pkg.go.dev/io/ioutil#ReadDir		deprecated!!
https://pkg.go.dev/os#ReadDir

*/

func osReadDir(path string) []fs.DirEntry {

	//read directory file
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func fsReadDir(path string) []fs.DirEntry {

	//read directory file
	files, err := fs.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func printFiles(files []fs.DirEntry) {
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
