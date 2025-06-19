package controllers

import (
    "net/http"
    "go-crud-api/models"
    "go-crud-api/utils"

    "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
    var creds models.User
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    for _, u := range models.Users {
        if u.Username == creds.Username && u.Password == creds.Password {
            token, _ := utils.GenerateToken(u.Username, u.Role)
            c.JSON(http.StatusOK, gin.H{"token": token})
            return
        }
    }

    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
