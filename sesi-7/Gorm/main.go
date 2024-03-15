package main

import (
	"Gorm/database"
	"Gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	createUser("johndoe@gmail.com")
	// getUserById(1)
}

func createUser(email string) {
	db := database.GetDB()	
	
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("User created successfully!", User)
}

func getUserById(id int) {
	db := database.GetDB()
	user:= models.User{}

	err := db.First(&user,"id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("User not found!")
			return
		}
		print("Error getting user data:", err)
	}

	fmt.Printf("User Data: %+v\n", user)
}