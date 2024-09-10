package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
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

func Userlist(c *gin.Context){

	type datatoSend struct {
		UserID       uint  
		Email    string 
	}
	currentUser , exists := c.Get("currentUser")
	if !exists{
		c.JSON(http.StatusUnauthorized , gin.H{
			"error":"Unauthorized",
		})
	}

	var response []datatoSend
	fmt.Println(currentUser)
	var user []models.User
	if err := database.GORM_DB.Find(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
        return
    }

	// var response []datatoSend
    for _, i := range user {
		if i.IsActive {
			response = append(response, datatoSend{
				UserID: i.ID,
				Email:  i.Email,
			})
		}
    }

    c.JSON(http.StatusOK, response)

}



func Follow(c *gin.Context) {
	userIDParam := c.Param("id")
	var userID uint
	if id, err := strconv.ParseUint(userIDParam, 10, 32); err == nil {
		userID = uint(id)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	currentUserid, exists := c.Get("currentUserid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	
	fmt.Println(currentUserid , userID)

	 
	if currentUserid == userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You cannot follow yourself",
		})
		return
	}
	followData := models.Follow{
		UserID:       uint(currentUserid),
		FollowedUserID: userID,
	}
	if err := database.GORM_DB.Create(&followData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to follow the user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully followed the user",
	})
}
