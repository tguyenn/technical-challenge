package main

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"type:varchar(100)"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}