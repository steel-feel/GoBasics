package main

import (
	"fmt"
)

//shadowing of constant
const a string = "hello";

//iota example, enumerated constants
const (
	//ignoring the first value, 
	// and starting from 1
	_ = iota;
	mango 
	orange
	pineapple
	grapes
)


const (
	//enumerated expression
	// shifting 1 bytewise to right
	//for first 1 << shifted by 0
	// second time 1 << shifted by 1, that became 10
	// third time 1 << shifted by 2, that became 100
	isAdmin = 1 << iota
	isDeveloper
	isManager
	isEndUser

)


func main() {
	//constant should be initalized with compile time value, 
	// not runtime value
	const a = 10;
	var b int16 = 4;
	var c int32 = 7

	//constants type has been inferred based on the action

	fmt.Printf("%d, %T\n", a+b, a+b);
	 fmt.Printf("%d, %T\n", a+c, a+c);

	//below will not work, due to type mismatch
	// const d int = 12;
	// fmt.Printf("%d, %T\n", b+c, b+c);

	//IOTA example
	iotaExample();

	//IOTA implementation
	iotaroles();

}

func iotaExample(){
	fmt.Printf("%d, %d, %d, %d\n", mango, orange, pineapple, grapes);
}

func iotaroles(){
	var role byte = isAdmin | isDeveloper;
	fmt.Printf("%b \n", role);

	// fmt.Printf("%b \n", isManager);

	//Check if role is Manager
	fmt.Printf("Has Manager access: %v \n" , ( isManager & role == isManager) );


}

