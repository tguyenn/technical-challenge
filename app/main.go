package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// User represents a user in the database
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"type:varchar(100)"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}

func main() {
    // PostgreSQL DSN (Data Source Name)
    dsn := "host=postgres user=user password=password dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

    // Connect to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    fmt.Println("✅ Connected to PostgreSQL successfully!")

    // AutoMigrate to create/update table schema
    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal("❌ Failed to migrate database:", err)
    }

    fmt.Println("✅ Database migration successful!")

    // Create a new user from command-line input
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter user name: ")
    name, _ := reader.ReadString('\n')
    // name = name[:len(name)-1] // Remove the newline character

    fmt.Print("Enter user email: ")
    email, _ := reader.ReadString('\n')
    // email = email[:len(email)-1] // Remove the newline character

    fmt.Print("Enter user password: ")
    password, _ := reader.ReadString('\n')
    // password = password[:len(password)-1] // Remove the newline character

    newUser := User{Name: name, Email: email, Password: password}
    result := db.Create(&newUser)

    if result.Error != nil {
        log.Fatal("❌ Failed to create user:", result.Error)
    }
    fmt.Println("✅ User created successfully with ID:", newUser.ID)

    // Fetch and print all users
    var users []User
    db.Find(&users)
    fmt.Println("👥 List of users:", users)
}