package main

import (
	"experiment/variable/constant"
	"experiment/variable/simple"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("experiment/variable")

	simple.Test()
	constant.Test()
}
