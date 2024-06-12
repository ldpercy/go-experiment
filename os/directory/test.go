// should probably consolidate this into a general os/filesystem experiment along with file

package directory

import (
	"fmt"
	"os"
)

func Test(directory string) {
	fmt.Println("Test: OS Directory")
	//TestRead(directory)
	TestWalk(directory)

	fmt.Println(os.DirFS(directory))
}

/*
https://pkg.go.dev/io/fs#ReadDir
https://pkg.go.dev/io/ioutil#ReadDir		deprecated!!
https://pkg.go.dev/os#ReadDir

*/
