package controllers

import (
	// "fmt"
	"fmt"
	"net/http"
	"time"

	// "os/user"
	"strconv"

	// "gorm.io/gorm"
	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/aayushkhosla/Mini-Social-Network/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
var validate = validator.New()

func UpdateUser(c *gin.Context){
	userID, _ := c.Get("currentUserid")
	UserID := userID.(uint)
	 var input struct {
		FirstName      string            `validate:"required,min=2,max=50"`
		LastName       string            `validate:"required,min=2,max=50"`
		DateOfBirth    time.Time         `validate:"required"`
		Gender         models.Gender     `validate:"required,gendercheek"`
		MaritalStatus  models.MaritalStatus `validate:"required,maritalstatuscheek"`
		
	 }
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
	user.Password = "";
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

	var user []models.User
	if err := database.GORM_DB.Find(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
        return
    }
	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)
	var response []datatoSend
    for _, i := range user {
		if currId != i.ID{
			response = append(response, datatoSend{
				UserID: i.ID,
				Email:  i.Email,
			})
		}
    }

    c.JSON(http.StatusOK, response)

}

func FollowingList(c *gin.Context){
	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)

	var followList []models.User
	err := database.GORM_DB.
    Table("follows").
    Select("users.*").
    Joins("JOIN users ON users.id = follows.followed_user_id").
    Where("follows.user_id = ? AND follows.active = ?", currId, true).
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

func FollowersList(c *gin.Context){
	currentUserid, _ := c.Get("currentUserid")
	currId := currentUserid.(uint)

	var followList []models.User
	err := database.GORM_DB.
    Table("follows").
    Select("users.*").
    Joins("JOIN users ON users.id = follows.user_id").
    Where("follows.followed_user_id = ? AND follows.active = ?", currId, true).
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