package main

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// darabase handler functions

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) error {
    return db.Create(user).Error
}

// GetUserByID retrieves a user by ID
func GetUserByID(db *gorm.DB, id int) (*User, error) {
    var user User
    if err := db.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

// UpdateUser updates an existing user
func UpdateUser(db *gorm.DB, user *User) error {
    return db.Save(user).Error
}

// DeleteUser deletes a user by ID
func DeleteUser(db *gorm.DB, id int) error {
    return db.Delete(&User{}, id).Error
}

// GetAllUsers retrieves all users
func GetAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}


// sets up api endpoints
func setupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // Create
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := CreateUser(db, &user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
            return
        }
        c.JSON(http.StatusCreated, user)
    })

    // Read
    r.GET("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        user, err := GetUserByID(db, id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, user)
    })

    // Update
    r.PUT("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        user, err := GetUserByID(db, id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        if err := c.ShouldBindJSON(user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := UpdateUser(db, user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
            return
        }
        c.JSON(http.StatusOK, user)
    })

    // Delete
    r.DELETE("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        if err := DeleteUser(db, id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    })

    // Read All
    r.GET("/users", func(c *gin.Context) {
        users, err := GetAllUsers(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
            return
        }
        c.JSON(http.StatusOK, users)
    })

    return r
}