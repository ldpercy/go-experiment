/* Structs
 */

package main

import (
	"experiment/struct/composite"
	"experiment/struct/simple"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Experiment: Structs")

	simple.Test()
	// composite.Test()
	composite.TestPointers()

}
