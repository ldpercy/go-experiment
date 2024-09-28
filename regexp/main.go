/*
https://pkg.go.dev/regexp
https://pkg.go.dev/regexp/syntax
*/

package main

import (
	_ "experiment/regexp/newline"
	_ "experiment/regexp/submatch"
	"experiment/regexp/zerowidth"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// submatch.Test()
	// newline.Test()
	zerowidth.Test()

}
