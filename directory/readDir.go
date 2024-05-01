package main

import (
	"io/fs"
	"log"
	"os"
)

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
