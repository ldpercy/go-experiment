Flow control
============


* elses must be cuddled





Loops
-----

https://www.learn-golang.org/en/Loops

All Go loops are 'for' loops, with variations on syntax to produce different kinds.

### Index loop:

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}


### While loop:

	for i < 10 {
		fmt.Println(i)
		i++
	}

### Range loop:

Loop over an array, slice or few other datastructures.

	for index, value := range myList {
		fmt.Printf("%d is index, %d is value", index, value)
	}

break, continue