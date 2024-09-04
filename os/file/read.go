package file

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func Test() {

	filename := "./file/test.txt"

	//open file
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	// output some info
	fmt.Println(reflect.TypeOf(content)) // content is a slice
	fmt.Println(file.Stat())

	fmt.Println(string(content))

}

func FileRead(filepath string) string {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error opening file: %v \n", err)
	}
	defer file.Close()

	log.Printf("File opened: %v \n", filepath)

	// Read file
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Error reading file: %v \n", err)
	}
	log.Printf("File read: %v \n", filepath)

	return string(content)
}
