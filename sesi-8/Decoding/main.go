package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func main() {
	var jsonString = `[{"full_name": "Airell", "email": "wWQp0@example.com", "age": 23}, {"full_name": "Siti", "email": "qE8kS@example.com", "age": 21}]`

	// var result Employee

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("result %d: %+v\n", i+1, v)
	}

	// fmt.Println("full_name", result.FullName)
	// fmt.Println("email", result.Email)
	// fmt.Println("age", result.Age)

	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])
}