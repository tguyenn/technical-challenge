// setup api endpoint and execute CRUD operations
// returns JSON data

package main

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// database interaction helper functions

// create
func dbCreateUser(db *gorm.DB, user *User) error {
    return db.Create(user).Error
}

// read
func dbReadUser(db *gorm.DB, id int) (*User, error) {
    var user User
    if err := db.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

// update
func dbUpdateUser(db *gorm.DB, user *User) error {
    return db.Save(user).Error
}

// delete
func dbDeleteUser(db *gorm.DB, id int) (bool, error) {
    result := db.Delete(&User{}, id)
    if result.Error != nil {
        return false, result.Error
    }
    return result.RowsAffected > 0, nil
}

// dump data
func dbDumpUsers(db *gorm.DB) ([]User, error) {
    var users []User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}


// register api endpoints. Gin will run the required operation on appropriate HTTP request
func setupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // disable debug message spam
    gin.SetMode(gin.ReleaseMode)


    // create
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := dbCreateUser(db, &user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
            return
        }
        c.JSON(http.StatusCreated, user)
    })

    // read
    r.GET("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        user, err := dbReadUser(db, id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, user)
    })

    // update
    r.PUT("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
    
        var updatedData User
        if err := c.ShouldBindJSON(&updatedData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        // Use GORM's Updates method to update only the provided fields
        if err := db.Model(&User{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
            return
        }
    
        // Fetch the updated user to return in the response
        var user User
        if err := db.First(&user, id).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user"})
            return
        }
    
        c.JSON(http.StatusOK, user)
    })

    // delete
    r.DELETE("/users/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
            return
        }
        deleted, err := dbDeleteUser(db, id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
            return
        }
        if !deleted {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    })

    // dump database
    r.GET("/users", func(c *gin.Context) {
        users, err := dbDumpUsers(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
            return
        }
        c.JSON(http.StatusOK, users)
    })

    return r
}