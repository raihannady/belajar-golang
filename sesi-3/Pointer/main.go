package main

import "fmt"

func main(){
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println(firstNumber)
	fmt.Println(&firstNumber)
	fmt.Println(*secondNumber)
	fmt.Println(secondNumber)

	fmt.Println("===========================")

	*secondNumber = 6

	fmt.Println(firstNumber)
	fmt.Println(&firstNumber)
	fmt.Println(*secondNumber)
	fmt.Println(secondNumber)
}