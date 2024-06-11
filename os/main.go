package main

import (
	"experiment/os/directory"
	"experiment/os/file"
)

func main() {

	dir := "."

	file.Test()
	directory.Test(dir)
}
