/*
https://go.dev/blog/constants
*/

package constant

import "log"

// So apparently there's a difference between these two:
const a = 1
const b int = 2

// I think it has to do with what can be coerced and what can't - lets try some stuff

// Untyped numeric constants can't overflow:
// const c int = 3298457389257423849047543509683450698 // overflow
const c = 3298457389257423849047543509683450698 // works

// You can construct strings as consts:
const (
	hello    = "hello"
	myName   = "Alice"
	greeting = hello + " " + myName
)

// What about regex?
const (
	prefix = `\(`
	middle = `[A-Za-z]+`
	suffix = `\)`
	exp    = prefix + middle + suffix
)

func Test() {
	log.Println("Constant Test")

	log.Println(a, b)

	log.Println(greeting)
	log.Println(exp)

	// log.Println(c) // overflows here though :(
}

/*
Examples of coercing untyped constants:
https://go.dev/play/p/2cgFoB4rYD


I think the principle is that untyped constants are a bit more flexible.

*/
