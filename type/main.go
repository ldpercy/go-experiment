/* Types

https://go.dev/ref/spec#Type_definitions

*/

package main

import (
	_ "experiment/type/alias"
	"experiment/type/check"
	_ "experiment/type/check"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Experiment: Types")
	check.Test()
	// alias.Test()
	// dynamic.Test()
}
