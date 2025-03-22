package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

var reader = bufio.NewReader(os.Stdin)

func createEntry(db *gorm.DB) {
	fmt.Println("Creating new entry!")
	fmt.Print("Enter user name: ")
    name, _ := reader.ReadString('\n')

    fmt.Print("Enter user email: ")
    email, _ := reader.ReadString('\n')

    fmt.Print("Enter user password: ")
    password, _ := reader.ReadString('\n')

    newUser := User{Name: name, Email: email, Password: password}
    result := db.Create(&newUser)

    if result.Error != nil {
        fmt.Println("❌ Failed to create user:", result.Error)
    } else {
		fmt.Println("Successfully created new user with ID: ", newUser.ID)
	}
}

func readEntry(db *gorm.DB) {
	fmt.Println("Reading entry!")
	fmt.Print("Enter User ID: ")
    userId, _ := reader.ReadString('\n')
	
	var user User
	result := db.First(&user, userId)

	if result.Error != nil {
        fmt.Println("❌ Failed to find user:", result.Error)
    } else {
		fmt.Println("Found user: ", user)
	}
}

func updateEntry(db *gorm.DB) {
	fmt.Println("Updating entry!")
	fmt.Print("Enter User ID: ")
    userId, _ := reader.ReadString('\n')
	
	var user User
	result := db.First(&user, userId)

	if result.Error != nil {
        fmt.Println("❌ Failed to find user:", result.Error)
    } else {
		fmt.Println("Found user: ", user)
	}
	
	fmt.Print("Enter new user name: ")
    user.Name, _ = reader.ReadString('\n')

    fmt.Print("Enter new user email: ")
    user.Email, _ = reader.ReadString('\n')

    fmt.Print("Enter new user password: ")
    user.Password, _ = reader.ReadString('\n')

    saveResult := db.Save(&user)
    if saveResult.Error != nil {
        fmt.Println("❌ Failed to update user:", saveResult.Error)
    } else {
        fmt.Println("✅ Successfully updated user with ID:", user.ID)
    }
}

func delEntry(db *gorm.DB) {
	fmt.Println("Deleting entry!")
	fmt.Print("Enter User ID: ")
    userId, _ := reader.ReadString('\n')
	
	var user User
	result := db.First(&user, userId)

	if result.Error != nil {
        fmt.Println("❌ Failed to find user:", result.Error)
    } else {
		fmt.Println("Deleting user: ", user)
	}

	db.Delete(&user)

}

func dumpData(db *gorm.DB) {
	fmt.Println("Dumping database")
    var users []User
    db.Find(&users)
    fmt.Println("List of users:", users)
}


func loopCLI(db *gorm.DB) {
    reader := bufio.NewReader(os.Stdin)
	for { // CLI while loop
		fmt.Println("What would you like to do? Please enter one of the follow keys and press enter: [C] Create [R] Read [U] Update [D] Delete [DD] Dump database [E] Exit")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		if action == "C" {
			createEntry(db)
		} else if action == "R" {
			readEntry(db)
		} else if action == "U" {
			updateEntry(db)
		} else if action == "D" {
			delEntry(db)
		} else if action == "DD" {
			dumpData(db)
		} else if action == "E" {
			fmt.Println("Exiting CLI. Goodbye!")
			break
		} else {
			fmt.Println("Invalid action :( try again")
		}
	}
}

func StartCLI(db *gorm.DB) {
	fmt.Println("Welcome to the greatest user management CLI of all time!")
	loopCLI(db)
}