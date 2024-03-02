package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var students = []map[string]string{
		{"name": "John", "address": "Cimahi", "jobs": "developer", "detail": "Golang is good for concurrency and performance."},
		{"name": "Alice", "address": "Cimahi", "jobs": "developer", "detail": "Golang provides a clean and efficient syntax."},
		{"name": "Bob", "address": "Cimahi", "jobs": "developer", "detail": "Prefers Golang for its simplicity and fast compilation."},
		{"name": "Charlie", "address": "Cimahi", "jobs": "developer", "detail": "Enjoys Golang's strong community support and documentation."},
		{"name": "David", "address": "Cimahi", "jobs": "developer", "detail": "Finds Golang suitable for building scalable and robust systems."},
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input:", os.Args[1])
		fmt.Println("Please enter a number")
		return
	}

	if index <= 0 || index >= len(students)+1 {
		fmt.Println("No Student found with no:", index)
		return
	}

	student := students[index-1]
	fmt.Printf("No: %d\nName: %s\nAddress: %s\nJobs: %s\nWhy Choose Golang?: %s\n", index, student["name"], student["address"], student["jobs"], student["detail"])
}
