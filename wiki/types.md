Types
=====

https://go.dev/ref/spec#Types

https://go.dev/ref/spec#Properties_of_types_and_values


	%T	a Go-syntax representation of the type of the value

Basic Format
------------

```go
	type newTypeName existingTypeConstructor
```

This creates a brand new, independent, type based on the existing type.



Type Aliasing
-------------

https://go.dev/ref/spec#Type_declarations

Aliasing requires the 'equals' symbol:

```go
	type number = int			// aliases an existing built-in type
	type Oblong = Rectangle		// aliases a user-defined type
```

Aliased types are entirely interchangeable with their source types.


Named Types
-----------

https://go.dev/ref/spec#Type_identity

Named types omit the 'equals' symbol:

```go
	type dollars int	// new type that 'copies' int, but is entirely independant
```

Named types are *not* substituable.



Conversion
----------

https://go.dev/ref/spec#Conversions




### Casting
https://stackoverflow.com/questions/19577423/how-to-cast-to-a-type-alias-in-go

> There is no casting in go. There are type assertions and type conversions.



Option types
------------

https://stackoverflow.com/questions/9993178/is-there-a-nice-way-to-simulate-a-maybe-or-option-type-in-go

```go
	Maybe X = *X
	Nothing = nil
	Just x = &x
```



Dynamic Types
-------------

https://stackoverflow.com/questions/55934210/creating-structs-programmatically-at-runtime-possible

https://medium.com/@utter_babbage/breaking-the-type-system-in-golang-aka-dynamic-types-8b86c35d897b

https://avilay.rocks/go-dynamic-types/
