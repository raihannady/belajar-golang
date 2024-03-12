package main

import "fmt"

type Employee struct {
	name string
	age int
	division string

}

// func main(){
// 	var employee1 Employee

// 	employee1.name = "Airell"
// 	employee1.age = 23
// 	employee1.division = "IT"

// 	fmt.Println(employee1.name)
// 	fmt.Println(employee1.age)
// 	fmt.Println(employee1.division)

// 	fmt.Println("========================")
	
// 	var employee2 Employee

// 	employee2.name = "Raihan"
// 	employee2.age = 23
// 	employee2.division = "Web Development"

// 	fmt.Println(employee2.name)
// 	fmt.Println(employee2.age)
// 	fmt.Println(employee2.division)
// }

func main(){
	var employee = []Employee{
		{name: "Airell", age: 23, division: "IT"},
		{name: "Raihan", age: 23, division: "Web Development"},
		{name: "Airell", age: 23, division: "IT"},
	}

	for _, v := range employee{
		fmt.Println("%+v/n", v)
	}
}