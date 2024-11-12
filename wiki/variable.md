Variables
=========

Types
-----

### var
A regular variable

### const
Constant, can only be simple values that can be evaluated by the compiler.

No constant maps, structs, arrays, slices


Constants
---------

https://go.dev/blog/constants



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