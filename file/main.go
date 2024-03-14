package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//read file
	//echo file

	filename := "test.txt"

	// show file stats
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Stat())

	// read the file
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))

	defer file.Close()
}
