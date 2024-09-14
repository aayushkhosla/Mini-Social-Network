package serialzer

import (
	"time"
	"github.com/aayushkhosla/Mini-Social-Network/models"
)


type Signupinput struct {
	Password       string            `validate:"required,min=6,max=50,passcheek"`
	Email          string            `validate:"required,email"`
	FirstName      string            `validate:"required,min=2,max=50"`
	LastName       string            `validate:"required,min=2,max=50"`
	DateOfBirth    time.Time         `validate:"required"`
	Gender         models.Gender     `validate:"required,gendercheek"`
	MaritalStatus  models.MaritalStatus `validate:"required,maritalstatuscheek"`
	EmployeeCode   string            `validate:"required,min=5,max=20"`
	OfficeAddress  string            `validate:"required,min=10,max=255"`
	OfficeCity     string            `validate:"required,max=50"`
	OfficeState    string            `validate:"required,max=50"`
	OfficeCountry  string            `validate:"required,max=50"`
	OfficeContactNo string           `validate:"required,min=10,max=15"`
	OfficeEmail    string            `validate:"required,email"`
	OfficeName     string            `validate:"required,min=3,max=100"`
	Address        string            `validate:"required,min=10,max=255"`
	City           string            `validate:"required,max=50"`
	State          string            `validate:"required,max=50"`
	Country        string            `validate:"required,max=50"`
	ContactNo1     string            `validate:"required,min=10,max=15"`
	ContactNo2     string            
}

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}


type UpdatePassword struct{
	Oldpassword string `validate:"required"`
	Newpassword string  `validate:"required"`
}
