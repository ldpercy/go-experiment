/*
https://pkg.go.dev/regexp/syntax
*/

package zerowidth

import (
	"log"
	"regexp"
)

func Test() {

	//str := "		foo : bar"

	//str := `		foo : bar`

	str := `this
		is
		a
		multiline
		string`

	reString := `(?m)(^)(.*)($)`

	re := regexp.MustCompile(reString)
	stringSubmatch := re.FindAllStringSubmatch(str, -1)

	log.Printf("%#v ", stringSubmatch)
}
