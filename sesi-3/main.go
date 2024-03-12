package main

import "belajar-golang/sesi-3/Exported"

func main(){
	Exported.Greet()

	var person = Exported.Person{}
	person.Invokegreet()
}