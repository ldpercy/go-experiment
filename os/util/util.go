package util

import (
	"fmt"
	"io/fs"
)

func PrintFiles(files []fs.DirEntry) {
	for _, file := range files {
		fmt.Printf("%#v \n", file)
	}
}
