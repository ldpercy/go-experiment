/*
https://pkg.go.dev/regexp
https://pkg.go.dev/regexp/syntax
*/

package main

import (
	_ "experiment/regex/newline"
	"experiment/regex/submatch"
	_ "experiment/regex/submatch"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	submatch.Test()
	// newline.Test()
	// zerowidth.Test()

}
