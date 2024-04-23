Data Structures
===============

* array
* slice
* map
* struct




Arrays
------

	// creating an array and assigning values later
	var studentsAge [10]int
	studentsAge = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// creating and assigning values to an array
	var studentsAge = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// creating and assigning values to an array without var keyword
	studentsAge := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// creating a nested array
	nestedArray := \[3\][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

Slices
------

	var sliceOfIntegers []int

	// creating slices from arrays
	fmt.Println(studentsAge[:4]) // [1 2 3 4]
	fmt.Println(studentsAge[6:]) // [7 8 9 10]
	fmt.Println(studentsAge[:])  // [1 2 3 4 5 6 7 8 9 10]

	// creating slices from slices
	secondSlice := firstSlice[1:5]
	fmt.Println(secondSlice) // [2 3 4 5]

	// creating a slice and assigning values later
	var tasksRemaining []string
	tasksRemaining = []string{"task 1", "task 2", "task 3"}

	// creating and assigning values to a slice
	var tasksRemaining = []string{"task 1", "task 2", "task 3"}

	// creating and assigning values to a slice without var keyword
	tasksRemaining := []string{"task 1", "task 2", "task 3"}


	// creating a nested slice
	nestedSlice := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}

Maps
----

	// creating a string -> int map
	var studentsAge map[string]int
	studentsAge = make(map[string]int)

	// creating a string -> int map
	studentsAge := make(map[string]int)

	// creating a map from literals
	studentsAge := map[string]int{
		"solomon": 19,
		"john":    20,
		"janet":   15,
		"daniel":  16,
		"mary":    18,
	}

	// creating nested maps
	studentResults := map[string]map[string]int{
		"solomon": {"maths": 80, "english": 70},
		"mary":    {"maths": 74, "english": 90},
	}

	// two-value assignment to get an existing key
	element, ok := studentsAge["solomon"]

	// two-value assignment to get a non-existing key
	element, ok = studentsAge["joel"]

Structs
-------

	type Rectangle struct {
		length  float64
		breadth float64
	}

	// creating a struct instance with var
	var myRectangle Rectangle

	// creating an empty struct instance
	myRectangle := Rectangle{}

	// creating a struct instance specifying values
	myRectangle := Rectangle{10, 5}

	// creating a struct instance specifying fields and values
	myRectangle := Rectangle{length: 10, breadth: 5}

	// you can also omit struct fields during their instantiation
	myRectangle := Rectangle{breadth: 10}


### Array and slice of struct


	arrayOfRectangles := [5]Rectangle{
		{10, 5},
		{15, 10},
		{20, 15},
		{25, 20},
		{30, 25},
	}


	sliceOfRectangles := []Rectangle{
		{10, 5},
		{15, 10},
		{20, 15},
		{25, 20},
		{30, 25},
	}


### Nested struct


	// creating a nested struct
	type address struct {
		houseNumber int
		streetName  string
		city        string
		state       string
		country     string
	}

	type Person struct {
		firstName   string
		lastName    string
		homeAddress address
	}

	// creating an instance of a nested struct
	person := Person{
		firstName: "Solomon",
		lastName:  "Ghost",
		homeAddress: address{
			houseNumber: 10,
			streetName:  "solomon ghost street",
			city:        "solomon city",
			state:       "solomon state",
			country:     "solomon country",
		},
	}

### Anonymous structs

	// creating a struct anonymously
	circle := struct {
		radius float64
		color  string
	}{
		radius: 10.6,
		color:  "green",
	}


### Struct Methods

	type Rectangle struct {
		length  float64
		breadth float64
	}

	func (r Rectangle) area() float64 {
		return r.length * r.breadth
	}

	// calling the Rectangle area method
	fmt.Println(myRectangle.area()) // 50



Advanced datastructures
-----------------------

https://github.com/Workiva/go-datastructures

	Augmented Tree
	Bitarray
	Futures
	Queue
	Fibonacci Heap
	Range Tree
	Set
	Threadsafe
	AVL Tree
	X-Fast Trie
	Y-Fast Trie
	Fast integer hashmap
	Skiplist
	Sort
	Numerics
	B+ Tree
	Immutable B Tree
	Ctrie
	Dtrie
	Persistent List
	Simple Graph









References
----------

https://blog.logrocket.com/comprehensive-guide-data-structures-go/