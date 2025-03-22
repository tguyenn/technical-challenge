package main

import (
	"fmt"
	"log"
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
	dsn := "host=localhost user=admin password=secret dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

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

	// Create a new user
	newUser := User{Name: "Toby", Email: "toby@example.com", Password: "secret123"}
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
