package main

import  "fmt"


func main(){

	//~~~~~~~
	//Boolean type
	//~~~~~~~~~~
	//Default is always 0
	var a bool = true

	fmt.Printf("Value: %v, Type: %T \n", a, a)

	//~~~~~~~
	//Integer type
	//~~~~~~~~~~
	var b int = 10
		c := 2 


	//Logic Operations
	fmt.Printf("& : %v \n", b & c)
	fmt.Printf("| : %v \n", b | c)
	fmt.Printf("^ : %v \n", b ^ c)
	fmt.Printf(" &^ : %v \n", b &^ c)

	//Bitwise operator
	//Its like (<<)multiplication and (>>)division with 2 ^ (second variable)
	fmt.Printf("<< : %v \n", b << c)
	fmt.Printf(">> : %v \n", b >> c)

	//~~~~~~~
	//Complex type
	//~~~~~~~~~~
	d := 2 + 2i
	var e complex64 = 2 + 4i
	//default is 128, implicit declaration
	f := complex(3, 5)

	fmt.Printf("Value: %v, Type: %T \n", d, d)
	fmt.Printf("Value: %v, Type: %T \n", e, e)
	fmt.Printf("Value: %v, Type: %T \n", f, f)

	//fetching real and imaginary part
	fmt.Printf("Real part: %v, Imaginary part: %v \n", real(f), imag(f))

	//~~~~~~~
	//String type
	//~~~~~~~~~~

	var g string = "Hello"
	 k := "World"

	//fetch single character, will fetch ascii value as its derievd from byte/int8
	fmt.Printf("Value: %v, Type: %T \n", g[0], g[0])
	//Fixing
	fmt.Printf("Value: %v, Type: %T \n", string(g[0]),string(g[0]))

	//concatenation
	fmt.Printf("Value: %v, Type: %T \n", g + k, g + k)

	// convert to byte stearm
	h := []byte(g);
	fmt.Printf("Value: %v, Type: %T \n", h,h)

	//~~~~~~~
	//Runes type, type alias for int32
	// good for UTF-32, functions like ReadRune(), in string package
	//~~~~~~~
	//notice the single quotes
	var i rune = 'a'
	j := 'b'

	fmt.Printf("Value: %v, Type: %T \n", i, i)
	fmt.Printf("Value: %v, Type: %T \n", j, j)


}