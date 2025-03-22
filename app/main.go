// Initialize database connection and enter into CLI loop

package main

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // PostgreSQL DSN (Data Source Name)
    dsn := "host=postgres user=user password=password dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // dont want to make global variable bc would also have to make this protected information global (bad)

    // Connect to the database
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

	fmt.Println("entering loop!")
	StartCLI(db)

}