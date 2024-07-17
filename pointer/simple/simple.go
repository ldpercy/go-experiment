package simple

import (
	"log"
)

func Test() {

	s := "a string"

	log.Println(s)

	ampersand_s := &s // &s is the address of s
	log.Println(ampersand_s)

	// asterisk_s := *s  // error - cannot get the underlying value of an ordinary var

	asterisk_ampersand_s := *ampersand_s // asterisk gets the value being pointed at the address

	log.Println(asterisk_ampersand_s)
}
