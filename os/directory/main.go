// should probably consolidate this into a general os/filesystem experiment along with file

package main

import (
	"experiment/directory/read"
	"experiment/directory/walk"
)

func main() {

	directory := "."

	read.Test(directory)

	walk.Test(directory)

}

/*
https://pkg.go.dev/io/fs#ReadDir
https://pkg.go.dev/io/ioutil#ReadDir		deprecated!!
https://pkg.go.dev/os#ReadDir

*/
