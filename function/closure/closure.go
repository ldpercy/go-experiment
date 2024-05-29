/*
Closure
*/

package closure

import "fmt"

func Test() {
	fmt.Println("Closures")
}

func example() {
	x := 10
	increase := func() {
		x = x + 1
	}
	increase()
	fmt.Println(x) // Outputs: 11
}
