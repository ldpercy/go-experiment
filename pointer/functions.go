package main

import "fmt"

func modify1(p person) {
	p.name = "Curly"
}

func modify2(p *person) {
	p.name = "Moe"
}

func print(name string, variable any) {
	fmt.Printf("%v: %#v \n", name, variable)
}
