package main

import ("fmt")

func main(){
	//pointers store only address, can be declared using *
	//When declaring pointer we are tell compiler 
	//what type of address the variable will hold

	var a int = 12;

	var b *int = &a;
	 
	fmt.Println(a,*b,&a,b);

	//go compliers help by mapping it to address 

	//function are pass by value, because they want it to be isolated env 
	// without affecting original data
	//pass by reference, pass by reference is when you want to change the value of the variable
	incrementValue(&a);

	fmt.Printf("Value after increment %v \n",a);
	
	
}


func incrementValue(c *int){
	*c = *c + 1;
	
	fmt.Printf("Pointer holding address %v \nAddress of pointer itself %v \n",c,&c);

}