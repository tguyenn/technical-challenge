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
func dbDeleteUser(db *gorm.DB, id int) error {
    return db.Delete(&User{}, id).Error
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
        user, err := dbReadUser(db, id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

    // buffer the user data so ID is preserved
    // todo: figure out if there is a better way to do this
    var updatedData User
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // update the fields of the existing user, but preserve ID
    user.Name = updatedData.Name
    user.Email = updatedData.Email
    user.Password = updatedData.Password

    // Update the user in the database
    if err := dbUpdateUser(db, user); err != nil {
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
        if err := dbDeleteUser(db, id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    })

    // Read All
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