/* Multiple Return
 */

package multireturn

import (
	"fmt"
)

func Test() {
	// Multiple Return
	fmt.Println(ReturnOne(5))
	//fmt.Println(multireturn.ReturnOneBracketed(6))
	fmt.Println(ReturnTwo(7))
	fmt.Println(ReturnThree(8))

	a, b := ReturnTwoFuncs()
	fmt.Println(a(11), b(12))
}

func ReturnOne(n int) int {
	return n * 2
}

/* brackets get removed by the formatter if there's only one return value
func ReturnOneBracketed(n int) (int) {
	return n * 2
}
*/

func ReturnTwo(n int) (int, int) {
	return n, n * 2
}

func ReturnThree(n int) (int, int, int) {
	return n, n * 2, n * 3
}

func ReturnTwoFuncs() (func(int) int, func(int) int) {

	funcA := func(a int) int { return 2 * a }
	funcB := func(b int) int { return 3 * b }

	return funcA, funcB

}
