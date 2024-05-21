Variables
=========

Types
-----

### var
A regular variable

### const
Constant, can only be simple values that can be evaluated by the compiler.

No constant maps, structs, arrays, slices






Blank Identifier - '_' (underscore)
-----------------------------------

https://go.dev/doc/effective_go#blank

> The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It's a bit like writing to the Unix /dev/null file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant.

