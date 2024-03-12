package main

import (
	"fmt"
	"strings"
)

// func main(){
// 	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

// 	fmt.Printf("%v", studentLists)
// }

// func print(names ...string) []map[string]string{
// 	var result []map[string]string

// 	for i, v := range names {
// 		key := fmt.Sprintf("student %d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

// func main(){
// 	numberLists := []int{1,2,3,4,5,6,7,8}

// 	result := sum(numberLists...)

// 	fmt.Println("Result", result)
// }

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, v := range numbers {
// 		total += v
// 	}
// 	return total
// }

func main() {
	profile("Airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func profile (name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("My favorite foods are", mergeFavFoods)
}
