
Arrays and Slices
=================

https://www.youtube.com/watch?v=Arb-LjPg7FA


Arrays are fixed length, slices are variable.
0 based

### Array

Arrays are fixed length once declared.

	var arrayName [size]arrayType = [size]arrayType{item1, item2, ...}

Or with type inference:

	var arrayName = [size]arrayType{item1, item2, ...}

Shorthand:

	arrayName := [size]arrayType{item1, item2, ...}


### Slice

Slices are variable length (use array type under the hood).


	var sliceName = []sliceType{item1, item2, ...}

Syntactic difference is the absence of number for size.

Can be appended to with:

	append(sliceName, value)

This creates a new slice, so to overwrite the old one:

	sliceName = append(sliceName, value)

### Slice ranges

	range := sliceName[start:end]	// start <= range < end

Open ended range:

	range := sliceName[start:]		// everything from start to end of array

Open start range

	range := sliceName[:end]			// everything before end


Ranges are slices and can be appended to:

	append(range, value)



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
