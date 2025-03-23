// manage user input and DB interaction

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"gorm.io/gorm"
)

var reader = bufio.NewReader(os.Stdin)

func createEntry(db *gorm.DB) {
	fmt.Print("Enter new user name: ")
    name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

    fmt.Print("Enter new user email: ")
    email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

    fmt.Print("Enter new user password: ")
    password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

    newUser := User{Name: name, Email: email, Password: password}
    result := db.Create(&newUser)

    if result.Error != nil {
        fmt.Println("Failed to create user with error", result.Error)
    } else {
		fmt.Println("Successfully created new user with ID:", newUser.ID)
	}
}

func readEntry(db *gorm.DB) {
	fmt.Print("Enter User ID to look up: ")
    input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid User ID! Please provide a number")
		return
	}
	
	var user User
	result := db.First(&user, userId)
	
	if result.Error != nil {
	fmt.Println("Failed to find user with error", result.Error)
	} else {
		fmt.Println("Found user: ", user)
	}
}
	
func updateEntry(db *gorm.DB) {
	fmt.Print("Enter User ID of User to update: ")

    input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid User ID! Please provide a number")
		return
	}
	
	var user User
	result := db.First(&user, userId)

	if result.Error != nil {
        fmt.Println("Failed to find user with error:", result.Error)
    } else {
		fmt.Println("Found user to update: ", user)
	}
	
	fmt.Print("Enter updated user name: ")
    user.Name, _ = reader.ReadString('\n')
	user.Name = strings.TrimSpace(user.Name)

    fmt.Print("Enter updated user email: ")
    user.Email, _ = reader.ReadString('\n')
	user.Email = strings.TrimSpace(user.Email)

    fmt.Print("Enter updated user password: ")
    user.Password, _ = reader.ReadString('\n')
	user.Password = strings.TrimSpace(user.Password)

    saveResult := db.Save(&user)
    if saveResult.Error != nil {
        fmt.Println("Failed to update user with error", saveResult.Error)
    } else {
        fmt.Println("Successfully updated user with ID:", user.ID)
    }
}

func delEntry(db *gorm.DB) {
	fmt.Print("Enter User ID to delete:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userId, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid User ID! Please provide a number")
		return
	}
	
	var user User
	result := db.First(&user, userId)

	if result.Error != nil {
        fmt.Println("Failed to find user with error", result.Error)
    } else {
		fmt.Println("Deleting user: ", user)
	}

	db.Delete(&user)

}

func dumpData(db *gorm.DB) {
    var users []User
    db.Find(&users)
    fmt.Println("List of users:")
	for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s\n", user.ID, user.Name, user.Email, user.Password)
    }
}


func loopCLI(db *gorm.DB) {
    reader := bufio.NewReader(os.Stdin)
	for { // CLI while loop
		fmt.Println("Please enter one of the following actions and press enter: [C] Create [R] Read [U] Update [D] Delete [DD] Dump database [E] Exit")
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
			fmt.Println("Exiting CLI and killing app container. Goodbye!")
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