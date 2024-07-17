Maps
----

https://www.youtube.com/watch?v=v3RodjH1i6c&list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&index=12

key:value pairs

All keys must be same type
All values must have same type

	mapName := map[keyType]valueType{
		key1: value1,
		key2: value2,
		...
	}

Reference an element:

	mapName[key]				// integer key
	mapName["stringKey"]

Looping:

	for k, v := range mapName {
		fmt.PrintLn(k, "-", v)
	}

Mutate:

	mapName[key]	= newValue
