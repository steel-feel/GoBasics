package main

import ("fmt" 
 "strconv")

//Grouping of variables
var (
//Package level variable for export
 I int = 10;
//Package level variable
 k float32 = 20;
)


func main(){

	//Local variable, default is float64
	j := 30.9;
	l := "40";


	//Printing package level variables
	fmt.Println(I);
	
	fmt.Println(j);
	fmt.Printf("%v , %T\n", I,I);
	fmt.Printf("%v , %T\n", k,k);
	fmt.Printf("%v , %T\n", j,j);
	fmt.Printf("%v , %T\n", l,l);

	//redeclaring variable, only for package level variables
	var k float64 = 30.9;

	fmt.Printf("%v , %T\n", k,k);

	//Converting float to int
	var m int = int(k);
	fmt.Printf("%v , %T\n", m,m);

	//Coverting int to string
	var n string = strconv.FormatFloat(k,'f',2,64);
	fmt.Printf("%v , %T\n", n,n);

	// Type conversion should be done explicitly if named type is different, 
	// stands true for struct as well
	TypeConversion()

}

func TypeConversion(){
	var a uint ;
	var b int ;

	//will cause issue if not explicitly converted
	// a = b;

	fmt.Printf("%v %v \n",a,b)
}


