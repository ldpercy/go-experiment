Defer panic recover
===================



Defer is fine


Panic & recover
---------------

Golang "Exception Handling" - ULTIMATE Golang Basics Tutorial
https://www.youtube.com/watch?v=3Km_P32ynIg

Not sure yet what the best way to use these are.
Probably not for simple errors/program flow - I expect these are better reserved for truly unusual circumstances.

Might be useful for goroutines?


### Panic Recovery in Go - Tutorial
https://www.youtube.com/watch?v=9iHQbcWmqUY

(!Good tutorial)

Last ditch mechanism built in to the language

Panic
	all code below the panic is not run
	deferred functions are run though
	panics propagate up the call stack

