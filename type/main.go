/* Types

https://go.dev/ref/spec#Type_definitions

*/

package main

import (
	_ "experiment/type/alias"
	_ "experiment/type/check"
	"experiment/type/dynamic"
	"fmt"
)

func main() {
	fmt.Println("Experiment: Types")
	// util.Test()
	// alias.Test()
	dynamic.Test()
}
