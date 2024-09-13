package controllers

import (

	"fmt"
	"net/http"
	"strconv"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/aayushkhosla/Mini-Social-Network/serialzer"
	"github.com/aayushkhosla/Mini-Social-Network/services"
	"github.com/aayushkhosla/Mini-Social-Network/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

)
var validate = validator.New()


func UpdatePassword(c *gin.Context){
	var passwordchangeinput serialzer.UpdatePassword

	if err := c.Bind(&passwordchangeinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	if err := validate.Struct(passwordchangeinput); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, err := range validationErrors {
			errorMessages[err.Field()] = fmt.Sprintf("The field %s is %s", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	services.UpdatePassword(c, passwordchangeinput)
}

func UpdateUser(c *gin.Context){
	
	 var input serialzer.UpdateInput;
	 if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate.RegisterValidation("gendercheek" , validation.Gendervalidation)
	validate.RegisterValidation("maritalstatuscheek" , validation.MaritalStatusvalidation)

	
	if err := validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, err := range validationErrors {
			errorMessages[err.Field()] = fmt.Sprintf("The field %s is %s", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}
	services.UpdateUser(c,input)
	
}

func Getuser(c *gin.Context) {
	userID, exists := c.Get("currentUserid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	UserID := userID.(uint)
	services.Getuser(c,UserID)
}

func Deleteuser(c *gin.Context){
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	Currentuser := currentUser.(models.User)
	services.Deleteuser(c,Currentuser)
}

func Userlist(c *gin.Context){

	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)

	services.Userlist(c,currId)

}

func FollowingList(c *gin.Context){
	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)
	services.FollowingList(c,currId)
}

func FollowersList(c *gin.Context){
	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)
	services.FollowersList(c,currId)
	
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
    currUserID := currentUserid.(uint)
    if currUserID == userID {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "You cannot follow yourself",
        })
        return
    }
	services.Follow(c,userID,currUserID)
	
}

func Unfollow(c *gin.Context){
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
	currUserID := currentUserid.(uint)

	if currUserID == userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You cannot Unfollow yourself",
		})
		return
	}
	services.Unfollow(c,userID,currUserID)
	
}