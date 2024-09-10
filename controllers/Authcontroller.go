package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {

	var logininput struct{
		Email string
		Password string
	}

	if err := c.Bind(&logininput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var userFound models.User
	fmt.Println(logininput)
	database.GORM_DB.First(  &userFound , "email=?", logininput.Email ).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invaild email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(logininput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or  password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}


func SignUp(c *gin.Context ){
		var input struct{
			Password       string         
			Username       string         
			Email          string               
			FirstName      string
			LastName       string
			DateOfBirth    time.Time
			Gender         models.Gender
			MaritalStatus  models.MaritalStatus
		}
		// if c.Bind(&input) != nil {
		// 	c.JSON(http.StatusBadRequest,gin.H{
		// 		"error": err,
		// 	})
		// 	return
		// }
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userFound models.User

		database.GORM_DB.First(&userFound , "email=?" ,input.Email).Find(&userFound)
		if userFound.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already used"})
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		user := models.User{
			Password :string(passwordHash),   
			Username      :input.Username,
			Email :input.Email  ,         
			// FirstName    :input.FirstName,
			// LastName       :input.LastName,
			// DateOfBirth    :datatypes.Date(input.DateOfBirth),
			// Gender       :input.Gender,
			// MaritalStatus  :input.MaritalStatus,
			FirstName    :"Aayush",
			LastName       :"khosla",
			DateOfBirth    : datatypes.Date(time.Now()),
		// 	"Gender":"male",
		// 	"MaritalStatus":"married""),
			Gender:"male",
			MaritalStatus:"married",
		}
		database.GORM_DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"data":user})

}

