Golang
======

https://go.dev/
https://pkg.go.dev/std
https://gobyexample.com/


Style
-----
https://go.dev/doc/effective_go
https://go.dev/wiki/CodeReviewComments
https://google.github.io/styleguide/go/decisions



Basics
------

* Go is pass-by-value
* Pass by value (non-pointer):			string int float boolean array struct
* Pass by reference (pointer-wrapper):	slice map function
* Arrays are fixed length, slices are variable length
* Map keys and values must all be of the same type



Rob Pike
--------

### Go Proverbs with Rob Pike

https://www.youtube.com/watch?v=PAAkCSZUG1c

* Don't communicate by sharing memory, share memory by communicating.
* Concurrency is not parallelism.
* Channels orchestrate; mutexes serialize.
* The bigger the interface, the weaker the abstraction.
* Make the zero value useful.
* interface{} says nothing.
* Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
* A little copying is better than a little dependency.
* Syscall must always be guarded with build tags.
* Cgo must always be guarded with build tags.
* Cgo is not Go.
* With the unsafe package there are no guarantees.
* Clear is better than clever.
* Reflection is never clear.
* Errors are values.
* Don't just check errors, handle them gracefully.
* Design the architecture, name the components, document the details.
* Documentation is for users.


### What We Got Right, What We Got Wrong | GopherConAU 2023

https://www.youtube.com/watch?v=yE5Tpp2BSGw

Done right:

*	started with formal spec
*	multiple compilers
*	portability
*	compatibility guarantee
*	library
*	tools
*	gofmt

Concurrency

	async await
	coloured functions
	https://en.wikipedia.org/wiki/Communicating_sequential_processes

	concurrency vs parallelism
	https://www.youtube.com/watch?v=oV9rvDllKEg

Interfaces

	sort.Interface len less swap
	empty writer reader
	maps slices arrays channels
	late introduction of generics
	actually parametric polymorphism

The compiler

	done in c
	llvm

Project management

	quality first
	svn-perforce-hg-git

Package management

	splash talk
	plain string for import stmt

Documentation

	even simple examples
	go playground

Q&A

	telemetry opt-in
	exceptions
	no Go 2
	Wish: Integers being aribtrary precision from the start (no integer overflow)   !!!!!!
	https://en.wikipedia.org/wiki/APL_(programming_language)

	stdlib
		syslog is garbage
		os.out should have interfaces not files

	memory arenas



### Concurrency is not Parallelism by Rob Pike

https://www.youtube.com/watch?v=oV9rvDllKEg

https://go.dev/talks/2012/waza.slide#1

concurrent: computer io devices, eg keyboard mouse display
parallel: vector dot product

CAR (Tony) Hoare - Communicating sequential processes 1978

Concurrent composition/decomposition

Goroutines
	channels
	select
	closures

Concurrency is not parallelism.
Concurrency enables parallelism.
Concurrency makes parallelism (and scaling and everything else) easy.



Learning
--------

Go Tutorial (Golang) for Beginners
https://www.youtube.com/watch?v=etSN4X_fCnM&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM






Pass by value
-------------

https://www.youtube.com/watch?v=LBLN4Wu5U4w&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=13

Group A: Non-pointer values
	string
	int
	float
	boolean
	array
	struct

	These are pass-by-value (copied)

	Pointers themselves are probably in this group

Group B: Pointer-wrapper values
	slice
	map
	function

	These effectively get passed by reference


Pointers
--------

Pointer to variable:

	valueName := "foobar"

	myPointer := &valueName 			// is pointer to valueName

	// 0x0123456789ab

Dereference pointer to get underlying value at its memory address:

	underlyingValue := *myPointer

Function that accepts a pointer:

	func updateValue(p *valueType) {
		*p := someNewValue
	}

(This explicit pointer syntax is implicit for pointer-wrapper types.)





