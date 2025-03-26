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

const baseURL = "http://localhost:8080" 
var reader = bufio.NewReader(os.Stdin)

// print the API responses in a more easily readable format
func prettyPrintJSON(jsonData []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonData, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(prettyJSON.String())
}

func createUser() {

	var user User

	fmt.Print("Enter new user name: ")
    user.Name, _ = reader.ReadString('\n')
	user.Name = strings.TrimSpace(user.Name) // delete trailing whitespace (newline chars and spaces)

    fmt.Print("Enter new user email: ")
    user.Email, _ = reader.ReadString('\n')
	user.Email = strings.TrimSpace(user.Email)

    fmt.Print("Enter new user password: ")
    user.Password, _ = reader.ReadString('\n')
	user.Password = strings.TrimSpace(user.Password)

    jsonData, _ := json.Marshal(user) // converts the user into a JSON data string

    resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonData)) // url, content type, actual content
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
	prettyPrintJSON(body)
}

func readUserByID() {

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
	prettyPrintJSON(body)
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

    req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/users/%s", baseURL, id), bytes.NewBuffer(jsonData)) // not natively supported in new/http package
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    // need this for non-natively supported HTTP requests
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
	prettyPrintJSON(body)
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
	prettyPrintJSON(body)
}

func readAllUsers() {
    resp, err := http.Get(baseURL + "/users")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
	prettyPrintJSON(body)
}


func loopCLI() {
    reader := bufio.NewReader(os.Stdin)
	for { // "while" loop
		fmt.Println("Please enter one of the following actions and press enter: [C] Create [R] Read [U] Update [D] Delete [DD] Dump database [E] Exit")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		if action == "C" {
			createUser()
		} else if action == "R" {
			readUserByID()
		} else if action == "U" {
			updateUser()
		} else if action == "D" {
			deleteUser()
		} else if action == "DD" {
			readAllUsers()
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