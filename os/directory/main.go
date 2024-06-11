// should probably consolidate this into a general os/filesystem experiment along with file

package directory

func Test(directory string) {
	TestRead(directory)
	TestWalk(directory)
}

/*
https://pkg.go.dev/io/fs#ReadDir
https://pkg.go.dev/io/ioutil#ReadDir		deprecated!!
https://pkg.go.dev/os#ReadDir

*/
