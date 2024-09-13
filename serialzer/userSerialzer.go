package serialzer

import (
	"time"
	"github.com/aayushkhosla/Mini-Social-Network/models"
)

type UpdateInput struct {
	FirstName      string            `validate:"required,min=2,max=50"`
	LastName       string            `validate:"required,min=2,max=50"`
	DateOfBirth    time.Time         `validate:"required"`
	Gender         models.Gender     `validate:"required,gendercheek"`
	MaritalStatus  models.MaritalStatus `validate:"required,maritalstatuscheek"`
	
 }

 type UserListFormate struct{
	UserID       uint  
	Email    string 
 }