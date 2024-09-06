/*
can i use basic string splits for newlines in a document?

*/

package newline

import (
	osfile "experiment/os/file"
	"log"
	"strings"
)

func Test() {

	/*
		txtCR := osfile.FileRead("./newline/CR.txt.bin")
		txtCRLF := osfile.FileRead("./newline/CRLF.txt.bin")
		txtLF := osfile.FileRead("./newline/LF.txt.bin")
	*/
	nlSrc := osfile.FileRead("./newline/newline.txt")

	splits := strings.Split(nlSrc, "\n") // this splits at proper newlines
	//splits := strings.Split(nlSrc, `\n`) // this splits at literal \n strings

	//log.Println(splits)
	log.Printf("%#v ", splits)
}
