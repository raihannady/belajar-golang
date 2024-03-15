package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)


type Employee struct {
	ID int
	Name string
	Age int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Raihan", Age: 23, Division: "Web Development"},
	{ID: 3, Name: "Badri", Age: 23, Division: "IT"},
}
var PORT = ":8080"

func main(){
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployees)
	fmt.Println("Server running on port", PORT)

	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET"{
		tpl, err := template.ParseFiles("index.html")

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, employees)
		// json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid request method", http.StatusBadRequest)
}

func createEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")
		convertAge, err := strconv.Atoi(age)
	
		if err != nil{
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
	
		newEmployee := Employee{
			ID: len(employees) + 1,
			Name: name,
			Age: convertAge,
			Division: division,
		}
	
		employees = append(employees, newEmployee)
	
		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid request method", http.StatusBadRequest)
}