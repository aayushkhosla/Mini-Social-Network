package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
)
var validate = validator.New()
func UpdatePassword(c *gin.Context){
	type input struct{
		Old_password string `validate:"required"`
		New_password string  `validate:"required"`
	}
	passwordchangeinput := input{}
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

	currentUser, exists := c.Get("currentUser")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    } 
		CurrentUser := currentUser.(models.User)
	if err := bcrypt.CompareHashAndPassword([]byte(CurrentUser.Password), []byte(passwordchangeinput.Old_password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Old password"})
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordchangeinput.New_password), bcrypt.DefaultCost)
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
func Login(c *gin.Context) {
	

	var loginInput struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	if err := c.Bind(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	if err := validate.Struct(loginInput); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, err := range validationErrors {
			errorMessages[err.Field()] = fmt.Sprintf("The field %s is %s", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	var userFound models.User
	if err := database.GORM_DB.First(&userFound, "email = ?", loginInput.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password."})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}	

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func SignUp(c *gin.Context ){
			var input struct{
				Password       string  `validate:"required"`        
				Email          string      `validate:"required,email"`
				FirstName      string	`validate:"required"`
				LastName       string `validate:"required"`
				DateOfBirth   datatypes.Date `validate:"required,date"`
				Gender         models.Gender `validate:"required"`
				MaritalStatus  models.MaritalStatus `validate:"required"`
				//office 
				EmployeeCode  string         `validate:"required"`
				OfficeAddress       string         `validate:"required"`
				OfficeCity          string `validate:"required,max=50"`
				OfficeState         string `validate:"required"`
				OfficeCountry       string `validate:"required"`
				OfficeContactNo     string `validate:"required"`
				OfficeEmail   string   `validate:"required,email"`
				OfficeName    string `validate:"required"`
				//personal 
				Address      string  `validate:"required,max=50"`
				City         string  `validate:"required"`
				State        string  `validate:"required"`
				Country      string  `validate:"required"`
				ContactNo1   string  `validate:"required"`
				ContactNo2   string  `validate:"required"`
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
			Email :input.Email  ,         
			FirstName    :input.FirstName,
			LastName       :input.LastName,
			DateOfBirth    :datatypes.Date(input.DateOfBirth),
			Gender       :input.Gender,
			MaritalStatus  :input.MaritalStatus,	
		}

		office := models.OfficeDetail{
			EmployeeCode :input.EmployeeCode,	
			Address       :input.OfficeAddress,
			City          :input.OfficeCity,
			State         :input.OfficeState,
			Country       :input.OfficeCountry,
			ContactNo     :input.OfficeContactNo,
			OfficeEmail   :input.OfficeEmail,
			OfficeName  :input.OfficeName,
		}

		adderss := models.AddressDetail{
			Address    :input.Address  ,
			City         :input.City,
			State        :input.State,
			Country      :input.Country,
			ContactNo1   :input.ContactNo1,
			ContactNo2   :input.ContactNo2,
		}
		user.AddressDetail = append(user.AddressDetail , adderss)
		user.OfficeDetail = append(user.OfficeDetail, office)

		database.GORM_DB.Create(&office)
		database.GORM_DB.Create(&adderss)

		database.GORM_DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"data":user})

}




