package main

import (
	"fmt"
)

func main() {

	//~~~~~~ if-else block ~~~~~~
	//	ifandElse();

	//~~~~~~~~~ switch ~~~~~~~~~~
	// switchBlock()

	//~~~~~~ for loop ~~~~~~
	//	forLoop();

	//~~~~~~ additional control flow ~~~~~~
	additionalControlFlow()

}

func ifandElse() {
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
}

func switchBlock() {
	//with initializer syntax, variable scope limited, mutliple cases can be clumped together
	switch c := 1; c {
	case 1, 2, 3:
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
	//~~~~~fallthrough and breaks~~~~~
	fallthroughbreaks()

	// type switches, for checking type of variable
	typeSwitches()
}

func fallthroughbreaks() {
	num := 2
	switch {
	case num > 0:
		fmt.Println("num is positive")
		//will fall through to next case, without checking the condition
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

func typeSwitches() {
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

func forLoop() {
	//for loop like javascript, also break and continue
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//like while loop
	for {
		fmt.Println("infinite loop aka while loop")
		break
	}

	//Labels
StartHere:

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break StartHere
		}
	}

	//Loops for collections
	students := [...]string{"john", "jane", "joe", "jasmine"}
	for i, name := range students {
		fmt.Println(i, name)
	}

}

func additionalControlFlow() {
	//defer
		deferFlow();

	//fmt.Print("should print last \n")
	//panic and recover
		panicRecover();

	//panic and recover simple example
	//paincAdvaned(3000)
	fmt.Print("this will print, because panic has been handled by recover \n")

}

func deferFlow() {
	a := 10
	//defer is executed in LIFO order after function execution finished, so the first defer is executed last
	//a will be set according to defer is pushed onto the stack
	defer fmt.Println("first defer, print last", a)

	fmt.Println("hello, world")

	a = 20
	defer fmt.Println("second defer, print first", a)

}

func panicRecover() {
	defer func() {

		//recover is a built-in function that is used to handle the panic	fmt.Println("should not print")
		if r := recover(); r != nil {
			fmt.Println("Recovered from ", r)
		}
	}()

	//panic is a built-in function that stops the normal flow of the program and returns control to the deferred function
	panic("something bad happened")

	//nothing will excute after panic
}

func paincAdvaned(port int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from ", r)
			port = port + 1
			paincAdvaned(port)
		}
	}()

	if port == 3000 {
		panic("port already in use")
	}

}
