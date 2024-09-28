/*
https://pkg.go.dev/regexp
https://pkg.go.dev/regexp/syntax
*/

package main

import (
	_ "experiment/regex/newline"
	_ "experiment/regex/submatch"
	"experiment/regex/zerowidth"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// submatch.Test()
	// newline.Test()
	zerowidth.Test()

}
