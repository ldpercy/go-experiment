package submatch

import (
	"log"
	"regexp"
)

func Test() {
	reString := `(\s*)(\w*)\s*(:?)\s*(.*)`
	str := "		foo : bar"

	re := regexp.MustCompile(reString)
	//testSubmatches(str, *re)

	fs := FindSubmatches(str, *re)
	log.Printf("%#v ", fs)
}

func testSubmatches(string string, regex regexp.Regexp) {

	stringSubmatch := regex.FindStringSubmatch(string)
	log.Printf("%#v ", stringSubmatch)

	stringIndex := regex.FindStringIndex(string)
	log.Printf("%#v ", stringIndex)

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)
	log.Printf("%#v ", stringSubmatchIndex)

}

/*
What I want is a function that returns an array of the submatches (only, not the whole string)
* their start and end char
* and the submatched strings
*/

type submatch struct {
	submatch string
	start    int
	end      int
}

func FindSubmatches(string string, regex regexp.Regexp) []submatch {
	result := []submatch{}

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)

	for i := 0; i < len(stringSubmatchIndex); i = i + 2 {
		result = append(result,
			submatch{
				submatch: string[stringSubmatchIndex[i]:stringSubmatchIndex[i+1]],
				start:    stringSubmatchIndex[i],
				end:      stringSubmatchIndex[i+1],
			},
		)
	}

	return result
}
