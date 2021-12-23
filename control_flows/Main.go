package main

import (
	"fmt"
)

func main() {
	a := 10

	//~~~~~~~~~ if else if ~~~~~~~~~~
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else if a < 10 {
		fmt.Println("a is less than 10")
	} else {
		fmt.Println("a is less than 5 and greater than 10")
	}

	//with intialization
	if b := 10; b > 5 {
		fmt.Println("b is greater than 5")
	} else {
		fmt.Println("b is less than 5")
	}

	//OR operator can short circuit
	if a > 5 || a < 10 {
		//second condition is not executed
		fmt.Println("a is greater than 5 or less than 10")
	}


	//~~~~~~~~~ switch ~~~~~~~~~~
	

	//with initializer syntax, variable scope limited, mutliple cases can be clumped together
	switch	c := 1 ; c {
	case 1,2,3:
		fmt.Println("a is 1 or 2 or 3")
	case 4, 5, 6:
		fmt.Println("a is 4 or 5 or 6")
	default:
		fmt.Println("a is not 5 or 10")
	}

	//switch without tag, cases need to be evaluated true 
	d := "banana"
	
	switch {
	case d == "apple":
		fmt.Println("fruit is apple")
	case d == "banana":
		fmt.Println("fruit is banana")
	default:
		fmt.Println("fruit is neither apple nor banana")
	}

	//fallthrough and breaks
	fallthroughbreaks();
	
	//type switches, for checking type of variable
	typeSwitches();

}

func fallthroughbreaks(){
	num := 2
	switch {
	case num > 0:
		fmt.Println("num is positive")
		fallthrough	
	case num > 1:
		if num < 5 {
			fmt.Println("num is less than 5")
			break
		} 	
		fmt.Println("num is less than 10")
	default:
		fmt.Println("num is negative")
	}
}

func typeSwitches(){

	var x interface{} = "hello"
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is of a different type")
	}
}
