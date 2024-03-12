package main

import (
	"fmt"
	"strings"
)

func main(){
	var names = []string{"Airell", "Jordan"}
	var printMsg = greet("Heii", names)

	fmt.Println(printMsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	return joinStr
}