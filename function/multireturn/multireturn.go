/* Multiple return
 */

package multireturn

import (
	"fmt"
)

func Test() {
	fmt.Println("Multiple return")
	// Multiple return
	fmt.Println(returnOne(5))
	//fmt.Println(multireturn.returnOneBracketed(6))
	fmt.Println(returnTwo(7))
	fmt.Println(returnThree(8))

	a, b := returnTwoFuncs()
	fmt.Println(a(11), b(12))
}

func returnOne(n int) int {
	return n * 2
}

/* brackets get removed by the formatter if there's only one return value
func returnOneBracketed(n int) (int) {
	return n * 2
}
*/

func returnTwo(n int) (int, int) {
	return n, n * 2
}

func returnThree(n int) (int, int, int) {
	return n, n * 2, n * 3
}

func returnTwoFuncs() (func(int) int, func(int) int) {

	funcA := func(a int) int { return 2 * a }
	funcB := func(b int) int { return 3 * b }

	return funcA, funcB
}
