package main

import (
	"fmt"
)

func modifyString1(s string) string {
	return fmt.Sprintf("%v %v", s, s)
}
