package read

import (
	"experiment/directory/util"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func Test(directory string) {

	// readDir
	fmt.Println("osReadDir")
	util.PrintFiles(osReadDir(directory))

	//fmt.Println("fsReadDir")
	//printFiles(fsReadDir(directory))

}

func osReadDir(path string) []fs.DirEntry {

	//read directory file
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

/*
func fsReadDir(path string) []fs.DirEntry {

	//read directory file
	files, err := fs.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
*/
