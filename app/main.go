// Initialize database connection and enter into CLI loop

package main

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    dsn := "host=postgres user=user password=password dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // global variable creds bad?
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("Connected to PostgreSQL successfully!")

    err = db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal("Failed to migrate database with error", err)
        return;
    }
    fmt.Println("Database migration successful!")

    fmt.Println("Starting the database CLI...")

	StartCLI(db)

}