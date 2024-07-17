/* Structs
 */

package main

import (
	_ "experiment/struct/composite"
	"experiment/struct/method"
	_ "experiment/struct/simple"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Experiment: Structs")

	// simple.Test()
	// composite.Test()
	// composite.TestPointers()
	method.Test()

}
