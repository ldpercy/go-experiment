Functions
=========


https://go.dev/doc/codewalk/functions/


Anonymous functions
-------------------

https://en.wikipedia.org/wiki/Anonymous_function#Go


	foo := func(x int) int {
		return x * x
	}
	fmt.Println(foo(10))



	// defined
	func(parameters) {

		// body
	}

	// defined and executed
	func(parameters) {

		// body
	} (arguments)


Function Literals
-----------------

https://go.dev/ref/spec#Function_literals




Lambdas
-------
https://www.gyata.ai/golang/golang-lambda/
https://www.koderhq.com/tutorial/go/function/#lambda




Multiple return
---------------

https://golangdocs.com/multiple-return-values-in-golang-functions


	func vals() (int, int) {
		return 3, 7
	}

	a, b := vals()


Function typing
---------------

https://go.dev/ref/spec#Function_types


Declare a typename for a function signature:

	type typeName func(args...) returnType



https://go.dev/ref/spec#Type_identity

> * Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.

I'm not sure I think this is a *good thing*, in particular the param/result name bit, but that's how it is.