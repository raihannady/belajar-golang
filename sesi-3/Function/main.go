package main

import "fmt"

func main(){
	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

	fmt.Printf("%v", studentLists)
}

func print(names ...string) []map[string]string{
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student %d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}
