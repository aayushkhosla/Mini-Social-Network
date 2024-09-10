package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Getuser( c *gin.Context){
	currentUser, exists := c.Get("currentUser")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

	fmt.Print(currentUser)
	c.JSON(http.StatusOK,currentUser)
}