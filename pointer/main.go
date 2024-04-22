package main

func main() {

	p1 := person{
		name: "Larry",
		age:  33,
	}

	//fmt.Printf("p1: %#v \n", p1)

	print("p1", p1)

	modify1(p1)
	print("p1", p1)

	modify2(&p1)
	print("p1", p1)

}
