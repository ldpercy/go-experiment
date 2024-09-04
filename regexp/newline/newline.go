/*
https://en.wikipedia.org/wiki/Newline
*/

package newline

import (
	osfile "experiment/os/file"
	"log"
	"regexp"
)

func Test() {

	re := `(\r\n|\n|\r)`

	testFile("./newline/CR.txt.bin", re)
	testFile("./newline/CRLF.txt.bin", re)
	testFile("./newline/LF.txt.bin", re)

}

func testFile(filepath string, reString string) {

	src := osfile.FileRead(filepath)

	re := regexp.MustCompile(reString)

	split := re.Split(src, -1)

	log.Printf("%#v", split)

	//for _, value := range split {
	//		log.Println(value)
	//	}

}
