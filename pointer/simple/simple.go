package simple

import "fmt"

func Test() {

	s := "a string"

	fmt.Println(s)

	sp := &s // address of s
	fmt.Println(sp)

}
