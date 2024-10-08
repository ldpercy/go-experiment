/* Types

https://go.dev/ref/spec#Type_definitions

*/

package main

import (
	_ "experiment/type/alias"
	"experiment/type/dynamic"
	_ "experiment/type/util"
	"fmt"
)

func main() {
	fmt.Println("Experiment: Types")
	// util.Test()
	// alias.Test()
	dynamic.Test()
}
