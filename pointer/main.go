package main

import (
	"experiment/pointer/function"
	_ "experiment/pointer/simple"
	_ "experiment/pointer/structs"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	//structs.Test()
	//simple.Test()
	function.Test()

}
