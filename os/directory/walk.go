/* Directory Walk

Not sure which to use here?

https://pkg.go.dev/io/fs#WalkDir
	func WalkDir(fsys FS, root string, fn WalkDirFunc) error
	type WalkDirFunc func(path string, d DirEntry, err error) error

https://pkg.go.dev/path/filepath#WalkDir
	func WalkDir(root string, fn fs.WalkDirFunc) error
	type WalkFunc func(path string, info fs.FileInfo, err error) error

There might be a typo in the signature for filepath.WalkDir - the second argument should be fs.DirEntry I think.

I think I should avoid using io/fs for now as it seems to be built for other purposes.

See also:

	https://pkg.go.dev/os@go1.22.4#ReadDir
	https://engineering.kablamo.com.au/posts/2021/quick-comparison-between-go-file-walk-implementations/

Maybe just use ReadDir?

*/

package directory

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func TestWalk(directory string) {
	// walkDir
	fmt.Println("walkDir")
	getFiles(directory)
	//result := getFiles(directory)
	//printFiles(result)
}

/*
	getFiles

try to get all files from path and deeper, and return full path
maybe with a globbing pattern
*/
func getFiles(path string) []string {
	var result = []string{}
	err := filepath.WalkDir(path, getFilesWalkFunc)

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", path, err)
		//return
	} else {
		result = append(result, path)
	}
	return result
}

func getFilesWalkFunc(path string, info fs.DirEntry, err error) error {
	subDirToSkip := "skip"
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if info.IsDir() && info.Name() == subDirToSkip {
		fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		return filepath.SkipDir
	}
	fmt.Printf("visited file or dir: %q\n", path)
	return nil
}
