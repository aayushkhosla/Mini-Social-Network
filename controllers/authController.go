package controllers

import (
	"fmt"
	"net/http"
	"github.com/aayushkhosla/Mini-Social-Network/serialzer"
	"github.com/aayushkhosla/Mini-Social-Network/services"
	"github.com/aayushkhosla/Mini-Social-Network/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

)


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
	services.Login(c, loginInput);
	
}

func SignUp(c *gin.Context ){
			var input serialzer.Signupinput

			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		
			validate.RegisterValidation("passcheek", validation.StrongPassword )
			validate.RegisterValidation("gendercheek" , validation.Gendervalidation)
			validate.RegisterValidation("maritalstatuscheek" , validation.MaritalStatusvalidation)
	
			if err := validate.Struct(input); err != nil {
				validationErrors := err.(validator.ValidationErrors)
				errorMessages := make(map[string]string)
				for _, err := range validationErrors {
					errorMessages[err.Field()] = fmt.Sprintf("Wrong information in field %s", err.Field())
				}
				c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
				return
			}

			services.SignUp(c,input)

}


