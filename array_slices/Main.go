package main

import (
	"fmt"
)

func main() {

	//~~~~ Arrays ~~~~
	// Have to give fixed size of array at the time of declaration
	//var a = [5]int{1,2,3,4,5} ;
	//OR
	//var a [5]int
	//OR, compute via elements
	var a = [...]int{1, 2, 3, 4, 5}

	//when assigned to another variable, it gets copied
	var b = a
	b[0] = 7
	fmt.Printf("a %v\n", a)
	fmt.Printf("b %v\n", b)

	//~~~~ Slices ~~~~
	// Slices are like arrays but they are dynamic in size
	// Slices are like references to arrays
	 var sliceA []int = []int{1, 2, 3, 4, 5}
	// OR, make function, if you know size and capacity
	// sliceA := make([]int, 5,5)
	// OR, from existing array
	// sliceA := a[:]

	var sliceB = sliceA


	//chaining any copy will reflect in all the variables
	sliceB[0] = 9
	
	fmt.Printf("sliceB %v\n", sliceB)
	fmt.Printf("length %v\n", len(sliceB))
	fmt.Printf("capacity %v\n", cap(sliceB))


	
	//~~~~ Slice Operations ~~~~
	//slicing 
	//sliceA[start:end], inclusive of start and exclusive of end
	fmt.Printf("sliceA %v\n", sliceA)
	fmt.Printf("from index 1, to index 2 %v \n", sliceA[1:3])
	//concatenation, using spread operator
	fmt.Printf("concatenation %v \n", append(sliceA, sliceB...))
	//append
	sliceA = append(sliceA, 6)



}
