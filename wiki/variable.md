Variables
=========

Regular variables
-----------------

```go
	// a variable must have a value or a type
	// var a						// syntax error

	// uninitialised variable:
	var b int

	var name1 = "Alice"				// implictly typed
	var name2 string = "Bob"		// explicitly typed

	// Shorthand syntax - within a function:
	name3 := "Carol"				// implictly typed

	// Cannot include types in shorthand declarations:
	// name4 string := "Dave"		// syntax error
```

Variables cannot be untyped.


Constants
---------

Constants can only be simple values that can be evaluated by the compiler, ie no constant maps, structs, arrays, slices.

Constants may be untyped.


```go
const i = 7			// untyped constant
const j int = 9		// typed constant
```

Some arithmetic operations can be performed at compile time to generate constants:

```go
const five = 5
const fivePlusThree = five + 3

const i = 1 iota
```


### Untyped Constants

Untyped numeric constants don't overflow:

```go
// const c int = 3298457389257423849047543509683450698	// typed constant overflows
const c = 3298457389257423849047543509683450698 		// untyped works
```

Untyped constants can be coerced into compatible types:
https://go.dev/play/p/2cgFoB4rYD

* https://youtu.be/VAf1UCOlMtk?t=639
* https://go.dev/blog/constants
* https://go.dev/ref/spec#Constants
* https://blog.learngoprogramming.com/learn-golang-typed-untyped-constants-70b4df443b61

Blank Identifier - '_' (underscore)
-----------------------------------

https://go.dev/doc/effective_go#blank

> The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It's a bit like writing to the Unix /dev/null file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant.



Naming
------

Variable names must begin with lower or upper case letters, or an underscore.

### Leading (initial) underscores

https://stackoverflow.com/questions/29891677/when-to-use-leading-underscore-in-variable-names-in-go

https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_