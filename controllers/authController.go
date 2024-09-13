package controllers

import (
	"fmt"
	"net/http"
	"os"

	// "regexp"

	// "regexp"
	"time"

	serialzer "github.com/aayushkhosla/Mini-Social-Network/Serialzer"
	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/aayushkhosla/Mini-Social-Network/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)




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
func Login(c *gin.Context) {
	

	var loginInput serialzer.LoginInput

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
			var input serialzer.Signupinput

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		
			validate.RegisterValidation("passcheck", validation.StrongPassword )
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


			// var userFound models.User
			
			// database.GORM_DB.First(&userFound , "email=?" ,input.Email).Find(&userFound)
			// if userFound.ID != 0 {
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Email already used"})
			// 	return
			// }
			// passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
			// if err != nil {
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// 	return
			// }
		
			// user := models.User{
			// 	Password :string(passwordHash),   
			// 	Email :input.Email  ,         
			// 	FirstName    :input.FirstName,
			// 	LastName       :input.LastName,
			// 	Gender       :input.Gender,
			// 	MaritalStatus  :input.MaritalStatus,	
			// }

			// office := models.OfficeDetail{
			// 	EmployeeCode :input.EmployeeCode,	
			// 	Address       :input.OfficeAddress,
			// 	City          :input.OfficeCity,
			// 	State         :input.OfficeState,
			// 	Country       :input.OfficeCountry,
			// 	ContactNo     :input.OfficeContactNo,
			// 	OfficeEmail   :input.OfficeEmail,
			// 	OfficeName  :input.OfficeName,
			// }

			// adderss := models.AddressDetail{
			// 	Address    :input.Address  ,
			// 	City         :input.City,
			// 	State        :input.State,
			// 	Country      :input.Country,
			// 	ContactNo1   :input.ContactNo1,
			// 	ContactNo2   :input.ContactNo2,
			// }
			// user.AddressDetail = append(user.AddressDetail , adderss)
			// user.OfficeDetail = append(user.OfficeDetail, office)
			// database.GORM_DB.Create(&user)


			// generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			// 	"id":  user.ID,
			// 	"exp": time.Now().Add(time.Hour * 4).Unix(),
			// })
		
			// token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			// 	return
			// }	
		
			// Response := gin.H{
			// 	"id":            user.ID,
			// 	"user_id":       user.ID,
			// 	"email":         user.Email,
			// 	"first_name":    user.FirstName,
			// 	"last_name":     user.LastName,
			// 	"last_modified": user.UpdatedAt.Format(time.RFC3339),
			// 	"gender":        user.Gender,
			// 	"marital_status": user.MaritalStatus,
			// 	"date_of_birth": user.DateOfBirth.Format("2006-01-02"),
			// 	"AddressDetail": gin.H{
			// 			"user_id":     user.ID,
			// 			"address":     adderss.Address,
			// 			"city":        adderss.City,
			// 			"state":       adderss.State,
			// 			"country":     adderss.Country,
			// 			"contact_no_1": adderss.ContactNo1,
			// 			"contact_no_2": adderss.ContactNo2,
			// 	},
			// 	"OfficeDetail": gin.H{
			// 			"user_id":      user.ID,
			// 			"employee_code": office.EmployeeCode,
			// 			"address":      office.Address,
			// 			"city":         office.City,
			// 			"state":        office.State,
			// 			"country":      office.Country,
			// 			"contact_no":   office.ContactNo,
			// 			"email":        office.OfficeEmail,
			// 			"name":         office.OfficeName,
			// 	},
			// 	"token": gin.H{
			// 		"key":         token,
			// 		"expiry_time": time.Now().Add(time.Hour * 4).Unix(),
			// 	},
				
			// }
			// c.JSON(http.StatusOK, gin.H{"data":Response})

}


