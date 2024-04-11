package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	filename := "test.txt"

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
