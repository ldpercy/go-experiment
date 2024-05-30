Types
=====

https://go.dev/ref/spec#Types

https://go.dev/ref/spec#Properties_of_types_and_values




Aliasing
--------

https://go.dev/ref/spec#Type_declarations

Examples - NB aliasing requires 'equals' symbol:

	type number = int			// aliases an existing built-in type
	type Oblong = Rectangle		// aliases a user-defined type

Aliased types are entirely interchangeable with their source types.


Named Types
-----------

https://go.dev/ref/spec#Type_identity

Examples - NB named types omit equals symbol:

	type dollars int	// new type that 'copies' int, but is entirely independant

