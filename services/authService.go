package services

import (
	"net/http"
	"os"
	"time"

	"github.com/aayushkhosla/Mini-Social-Network/database"
	"github.com/aayushkhosla/Mini-Social-Network/models"
	"github.com/aayushkhosla/Mini-Social-Network/serialzer"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func SignUp(c *gin.Context , input serialzer.Signupinput){
	var userFound models.User
			
	database.GORM_DB.First(&userFound , "email=?" ,input.Email)

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
		DateOfBirth: input.DateOfBirth,         
		FirstName    :input.FirstName,
		LastName       :input.LastName,
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
	if err := database.GORM_DB.Create(&user).Error ; err!= nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : "Email already used"})
		return
	}


	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}	

	Response := gin.H{
		"id":            user.ID,
		"user_id":       user.ID,
		"email":         user.Email,
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"last_modified": user.UpdatedAt.Format(time.RFC3339),
		"gender":        user.Gender,
		"marital_status": user.MaritalStatus,
		"date_of_birth": user.DateOfBirth, 	
		"AddressDetail": gin.H{
				"user_id":     user.ID,
				"address":     adderss.Address,
				"city":        adderss.City,
				"state":       adderss.State,
				"country":     adderss.Country,
				"contact_no_1": adderss.ContactNo1,
				"contact_no_2": adderss.ContactNo2,
		},
		"OfficeDetail": gin.H{
				"user_id":      user.ID,
				"employee_code": office.EmployeeCode,
				"address":      office.Address,
				"city":         office.City,
				"state":        office.State,
				"country":      office.Country,
				"contact_no":   office.ContactNo,
				"email":        office.OfficeEmail,
				"name":         office.OfficeName,
		},
		"token": gin.H{
			"key":         token,
			"expiry_time": time.Now().Add(time.Hour * 4).Unix(),
		},
		
	}
	c.JSON(http.StatusOK, gin.H{"data":Response})

}


func Login(c *gin.Context , loginInput serialzer.LoginInput){
	var userFound models.User
	if err := database.GORM_DB.First(&userFound, "email = ?", loginInput.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password."})
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
	Response := gin.H{
		"token":gin.H{
			"key" : token,
			"expiry time" : time.Now().Add(time.Hour * 4).Unix(),
		},
		"user" :gin.H{
			"id":userFound.ID,
			"email" :userFound.Email,
			"first_name" : userFound.FirstName,
			"last_name" : userFound.LastName,
			"last_modified" : userFound.UpdatedAt,
		},
	}	
		
	c.JSON(http.StatusOK,Response )
}