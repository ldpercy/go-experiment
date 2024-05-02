package folder

import "fmt"

var PublicVar = "this is public (exported) var of the package"
var privateVar = "this variable is private to the package"

func FolderStuff() {
	fmt.Println(PublicVar)
	fmt.Println(privateVar)
}
