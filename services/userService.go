package services

import (
	"net/http"

	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/aayushkhosla/Mini-Social-Network/serialzer"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func UpdatePassword(c *gin.Context ,passwordchangeinput serialzer.UpdatePassword ){
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	} 
		CurrentUser := currentUser.(models.User)
	if err := bcrypt.CompareHashAndPassword([]byte(CurrentUser.Password), []byte(passwordchangeinput.Oldpassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Old password"})
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordchangeinput.Newpassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	CurrentUser.Password = string(passwordHash)
	if err:= database.GORM_DB.Save(&CurrentUser).Error ; err!=nil{
		c.JSON(http.StatusInternalServerError ,gin.H{
			"error":"Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK ,gin.H{
		"message" : "Operation successful",
	} )
}


func UpdateUser(c *gin.Context , input serialzer.UpdateInput){
	userID, _ := c.Get("currentUserid")
	UserID := userID.(uint)
	var userFound models.User	

	if err := database.GORM_DB.Preload("AddressDetail").Preload("OfficeDetail").Where("id = ?", UserID).First(&userFound).Error; err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	userFound.FirstName = input.FirstName
	userFound.LastName = input.LastName
	userFound.DateOfBirth  =input.DateOfBirth
	userFound.Gender = input.Gender
	userFound.MaritalStatus = input.MaritalStatus
	if err := database.GORM_DB.Save(&userFound).Error ; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK  , userFound)
}

func Getuser(c *gin.Context ,UserID uint){
	var user models.User
	if err := database.GORM_DB.Preload("AddressDetail").Preload("OfficeDetail").Where("id = ?", UserID).First(&user).Error; err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

func Deleteuser(c *gin.Context , Currentuser models.User){
	if err := database.GORM_DB.Delete(&Currentuser).Error; err!= nil{
		c.JSON(http.StatusInternalServerError , gin.H{
			"error":"Internal Server Error",
		})
	
	}
	if err := database.GORM_DB.Where("user_id = ? OR followed_user_id = ?" ,Currentuser.ID ,Currentuser.ID ).Delete(&models.Follow{}).Error ; err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{
			"error":"Internal Server Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func Userlist(c *gin.Context , currId uint){
	var user []models.User
	if err := database.GORM_DB.Find(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
        return
    }
	
	var response []serialzer.UserListFormate
    for _, i := range user {
		if currId != i.ID{
			response = append(response, serialzer.UserListFormate{
				UserID: i.ID,
				Email:  i.Email,
			})
		}
    }

    c.JSON(http.StatusOK, response)
}

func FollowingList(c *gin.Context , currID uint){
	var followList []models.User
	err := database.GORM_DB.
    Table("follows").
    Select("users.*").
    Joins("JOIN users ON users.id = follows.followed_user_id").
    Where("follows.user_id = ? AND follows.active = ?", currID, true).
    Find(&followList).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error retrieving follow list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list":followList,
	})
}

func FollowersList(c *gin.Context , currID uint){
	var followList []models.User
	err := database.GORM_DB.
    Table("follows").
    Select("users.*").
    Joins("JOIN users ON users.id = follows.user_id").
    Where("follows.followed_user_id = ? AND follows.active = ?", currID, true).
    Find(&followList).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error retrieving follow list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list":followList,
	})
}

func Follow(c *gin.Context , userID uint , currUserID uint ){
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


func Unfollow(c *gin.Context , userID uint , currUserID uint){
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
