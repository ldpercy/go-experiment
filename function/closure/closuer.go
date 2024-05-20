/*
Closure
*/

package closure

import "fmt"

func test() {
	fmt.Println("lambda main")
}

func example() {
	x := 10
	increase := func() {
		x = x + 1
	}
	increase()
	fmt.Println(x) // Outputs: 11
}
