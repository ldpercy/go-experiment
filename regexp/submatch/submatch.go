package submatch

import (
	"log"
	"regexp"
)

func Test() {
	reString := `(^)(\t*)(\w*)(\s*)(:)(\s*)(.*)($)`
	//str := "		foo : bar"
	str := "name : Type value" // 17 characters

	re := regexp.MustCompile(reString)
	testSubmatches(str, *re)

	//fs := FindSubmatches(str, *re)
	//log.Printf("%#v ", fs)

	submatchSpecial := FindSubmatchSpecial(str, *re)
	log.Printf("%v ", submatchSpecial)
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

type Submatch struct {
	Submatch string
	Start    int
	End      int
}

func FindSubmatches(string string, regex regexp.Regexp) []Submatch {
	result := []Submatch{}

	stringSubmatchIndex := regex.FindStringSubmatchIndex(string)

	for i := 0; i < len(stringSubmatchIndex); i = i + 2 {
		result = append(result,
			Submatch{
				Submatch: string[stringSubmatchIndex[i]:stringSubmatchIndex[i+1]],
				Start:    stringSubmatchIndex[i],
				End:      stringSubmatchIndex[i+1],
			},
		)
	}

	return result
}

func FindSubmatchesAll(string string, regex regexp.Regexp) []Submatch {
	result := []Submatch{}

	stringSubmatchIndex := regex.FindAllStringSubmatchIndex(string, -1)

	for _, value := range stringSubmatchIndex {
		for i := 0; i < len(value); i = i + 2 {
			result = append(result,
				Submatch{
					Submatch: string[value[i]:value[i+1]],
					Start:    value[i],
					End:      value[i+1],
				},
			)
		}
	}

	return result
}
