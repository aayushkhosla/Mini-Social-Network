package controllers

import (

	"fmt"
	"net/http"

	// "os/user"
	"strconv"

	// "gorm.io/gorm"
	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/gin-gonic/gin"
)

func updateUser(c *gin.Context){

}



func Getuser(c *gin.Context) {
	userID, exists := c.Get("currentUserid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	UserID := userID.(uint)
	var user models.User
	if err := database.GORM_DB.Preload("AddressDetail").Preload("OfficeDetail").Where("id = ?", UserID).First(&user).Error; err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, user)
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
	if err := database.GORM_DB.Delete(&Currentuser).Error; err!= nil{
			c.JSON(http.StatusInternalServerError , gin.H{
				"error":"Internal Server Error",
			})
		
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
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
    // for _, i := range user {
	// 	if i.IsActive {
	// 		response = append(response, datatoSend{
	// 			UserID: i.ID,
	// 			Email:  i.Email,
	// 		})
	// 	}
    // }

    c.JSON(http.StatusOK, response)

}

func FollowList(c *gin.Context){
	

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
	var userCheek models.User
	if err := database.GORM_DB.First(&userCheek,userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}


    var existingFollow models.Follow
     res := database.GORM_DB.Where("user_id = ? AND followed_user_id = ?", currUserID, userID).First(&existingFollow)

    if res.RowsAffected > 0 {
        if existingFollow.Active {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "You are already following this user.",
            })
            return
        }

        existingFollow.Active = true
        if err := database.GORM_DB.Save(&existingFollow).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to reactivate the follow",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Successfully reactivated the follow.",
        })
    } else {
        newFollow := models.Follow{
            UserID:        currUserID,
            FollowedUserID: userID,
            Active:        true,
        }

        if err := database.GORM_DB.Create(&newFollow).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to follow the user",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Successfully followed the user",
        })
    }
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
	var userCheek models.User
	if err := database.GORM_DB.First(&userCheek,userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	var mydata models.Follow
	database.GORM_DB.Where("user_id = ? AND followed_user_id = ?", currUserID, userID).First(&mydata)

	if mydata.ID == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "Follow relationship not found",
		})
		return
	}else{
		if mydata.Active {
			mydata.Active = false
			database.GORM_DB.Save(&mydata)
			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully Unfollowed the user",
			})
			return

		}else{
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "Follow relationship not found",
			})
			return

		}

	}

}