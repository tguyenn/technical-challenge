// Initialize database connection, init Gin router, and enter into CLI loop

package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {

    // fetch dynamic environment variables defined in docker-compose.yml
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")
    timezone := os.Getenv("DB_TIMEZONE")
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
        return;
    }
    fmt.Println("Connected to PostgreSQL successfully!")
    
    // flag to ensure CLI only starts after the Gin router is up and running
    routerReady := make(chan bool)

    // make goroutine that allows API server to not block the CLI from starting up
    go func() {
        r := setupRouter(db)
        fmt.Println("Gin router initialized!")
        routerReady <- true
        if err := r.Run(":8080"); err != nil {
            log.Fatal("Failed to start Gin server:", err)
        }
    }()
    
    <-routerReady // wait until API ready
    fmt.Println("Starting the database CLI...")

	loopCLI()

}