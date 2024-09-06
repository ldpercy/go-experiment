/*
https://en.wikipedia.org/wiki/Newline
*/

package newline

import (
	osfile "experiment/os/file"
	"experiment/regexp/submatch"
	"log"
	"regexp"
)

func Test() {

	txtCR := osfile.FileRead("./newline/CR.txt.bin")
	txtCRLF := osfile.FileRead("./newline/CRLF.txt.bin")
	txtLF := osfile.FileRead("./newline/LF.txt.bin")

	reSplit := regexp.MustCompile(`(\r\n|\n|\r)`)

	splitFile(txtCR, *reSplit)
	splitFile(txtCRLF, *reSplit)
	splitFile(txtLF, *reSplit)

	//reSubmatch := regexp.MustCompile(`([^\r\n]+)(\r\n|\n|\r)`)
	reSubmatch := regexp.MustCompile(`([^\r\n]+(\r\n|\n|\r))`)

	submatchFile(txtCR, *reSubmatch)
	submatchFile(txtCRLF, *reSubmatch)
	submatchFile(txtLF, *reSubmatch)

}

func splitFile(text string, re regexp.Regexp) {

	split := re.Split(text, -1)

	log.Printf("%#v", split)

	//for _, value := range split {
	//		log.Println(value)
	//	}

}

func submatchFile(text string, re regexp.Regexp) {

	submatch := submatch.FindSubmatches(text, re)

	log.Printf("%#v", submatch)

	//for _, value := range split {
	//		log.Println(value)
	//	}

}
