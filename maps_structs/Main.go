package main

import (
	"fmt"
	 "reflect"
)



	type Vehicle struct {
		name string
		wheels int
	
	}


func main() {
	//~~~~ Map ~~~~
	//Its pass by reference
	student := map[string]int{
		"Asha" : 56,
		"Raj" : 23,
		"Ravi" : 98,
	}

	student2 := student
	//Or, declare using make()
	//student := make(map[string]int)

	//Add elements to map
	student["Shyam"] = 32
	//accessing the map
	fmt.Println(student["Asha"])
	//delete key
	delete(student, "Asha")
	//check if key exists
	AshaMarks, ok := student["Asha"]
	fmt.Println(AshaMarks,ok)

	fmt.Println(student)
	fmt.Println(student2)

	
	//~~~~ Struct ~~~~
	//It is pass by value
	v1 := Vehicle{ name: "Ship", wheels : 0 }

	fmt.Println(v1)

	//Annonymous Struct
	v2 := struct { name string ; wheels int} { name: "Bike" ,
	 wheels : 2 ,
	}

	fmt.Println(v2)

	//Embedding Struct
	embedding()

	//Tags, just for developer to know what restrictions are there per field
	Tags()

}


func embedding(){

	//embedding struct
	type Car struct {
		Vehicle
		name string
		brand string
	}
	c := Car{}
	c.Vehicle.name = "Car"
	c.name = "S1"
	c.brand = "Tesla"
	c.wheels = 4
	

	d := Car{
		Vehicle: Vehicle{ name: "Car", wheels: 4},
		name: "Kwid",
		brand: "Renault",
	}
	

	fmt.Println(c,d)

}


func Tags() {
	type Person struct {
		Name string `max: 100 `
		Age int `required max: 150`
	}
	p := Person{
		Name: "Asha",
		Age: 230,
	}
	fmt.Println(p)

	t := reflect.TypeOf(p)
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}