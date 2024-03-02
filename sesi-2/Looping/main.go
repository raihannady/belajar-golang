package main

import "fmt"

// func main (){
// 	for i := 0; i < 5; i++ {
// 		for j := i; j < 5; j++ {
// 			fmt.Print(j,  " ")
// 		}
// 	}
// 	fmt.Println()
// }


func main(){
	outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke -", i+1)
		for j := 0; j < 3; j++ {
			if j == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}