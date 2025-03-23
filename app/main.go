// Initialize database connection and enter into CLI loop

package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {

    // fetch dynamic envionrment variables defined in docker-compose.yml
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")
    timezone := os.Getenv("DB_TIMEZONE")
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

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