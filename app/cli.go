// manage user input and makes appropriate HTTP requests to database

package main

import (
	"bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"

	"bufio"
	"fmt"
	"os"
	"strings"
)

const baseURL = "http://localhost:8080" // Replace with your API's base URL

var reader = bufio.NewReader(os.Stdin)

func createUser() {

	var user User

	fmt.Print("Enter updated user name: ")
    user.Name, _ = reader.ReadString('\n')
	user.Name = strings.TrimSpace(user.Name)

    fmt.Print("Enter updated user email: ")
    user.Email, _ = reader.ReadString('\n')
	user.Email = strings.TrimSpace(user.Email)

    fmt.Print("Enter updated user password: ")
    user.Password, _ = reader.ReadString('\n')
	user.Password = strings.TrimSpace(user.Password)

    jsonData, _ := json.Marshal(user)

    resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

func getUserByID() {

	fmt.Print("Enter User ID to lookup:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

    resp, err := http.Get(fmt.Sprintf("%s/users/%s", baseURL, id))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}
	
func updateUser() {
	
	var user User

	fmt.Print("Enter User ID of User to update: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Enter updated user name: ")
    user.Name, _ = reader.ReadString('\n')
	user.Name = strings.TrimSpace(user.Name)

    fmt.Print("Enter updated user email: ")
    user.Email, _ = reader.ReadString('\n')
	user.Email = strings.TrimSpace(user.Email)

    fmt.Print("Enter updated user password: ")
    user.Password, _ = reader.ReadString('\n')
	user.Password = strings.TrimSpace(user.Password)

    jsonData, _ := json.Marshal(user)

    req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/users/%s", baseURL, id), bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

func deleteUser() {

	fmt.Print("Enter User ID to delete:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

    req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/users/%s", baseURL, id), nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}

func getAllUsers() {
    resp, err := http.Get(baseURL + "/users")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}


func loopCLI() {
    reader := bufio.NewReader(os.Stdin)
	for { // CLI while loop
		fmt.Println("Please enter one of the following actions and press enter: [C] Create [R] Read [U] Update [D] Delete [DD] Dump database [E] Exit")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		if action == "C" {
			createUser()
		} else if action == "R" {
			getUserByID()
		} else if action == "U" {
			updateUser()
		} else if action == "D" {
			deleteUser()
		} else if action == "DD" {
			getAllUsers()
		} else if action == "E" {
			fmt.Println("Exiting CLI and killing app container. Goodbye!")
			break
		} else {
			fmt.Println("Invalid action :( try again")
		}
	}
}

func StartCLI() {
	fmt.Println("Welcome to the greatest user management CLI of all time!")
	loopCLI()
}