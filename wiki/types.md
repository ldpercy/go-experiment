Types
=====

```go
	type wholeNumber = int			// type alias
	type myStruct struct			// named type
	fmt.Sprintf("%T", variable)		// type of variable
```

https://go.dev/ref/spec#Types

https://go.dev/ref/spec#Properties_of_types_and_values


Type Aliases
------------

Aliases are synonyms and entirely interchangeable with their source types.

Aliasing requires the `=` symbol:

```go
	type wholeNumber = int		// aliases an existing built-in type
	type Oblong = Rectangle		// aliases a user-defined type
```

https://go.dev/ref/spec#Type_declarations


Named Types
-----------

Named types are independent of their sources and *not* substitutable.

Named types omit the `=` symbol:

```go
	type dollars int	// new type that 'copies' int, but is entirely independant
```


https://go.dev/ref/spec#Type_identity


Type checking
-------------

String formatting:
```go
	fmt.Sprintf("%T", variable)
```

Reflection:
```go
	reflect.TypeOf(variable)
```

Reflect kind - returns underlying type of named types:
```go
	reflect.ValueOf(variable).Kind()
```


https://golangtutorial.dev/tips/find-type-of-an-object/
https://golangtutorial.dev/tips/find-type-of-an-object/


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
