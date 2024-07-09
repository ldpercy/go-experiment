package main

import (
	"fmt"
	"regexp"
)

func main() {

	re := regexp.MustCompile(`(\s*)(\w*)\s*(:?)\s*(.*)`)

	str := "		foo : bar"

	result := re.FindStringSubmatch(str)

	fmt.Printf("%#v ", result)

}
