package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Gender string

const (
    Male   Gender = "male"
    Female Gender = "female"
    Other  Gender = "other"
)

type MaritalStatus string
const (
    Single  MaritalStatus = "single"
    Married MaritalStatus = "married"
)

type User struct {
    gorm.Model
    ID             uint           `gorm:"primaryKey"`
    Password       string         `gorm:"not null"`
    Username       string         `gorm:"not null"`
    Email          string         `gorm:"unique;not null"`
    IsActive       bool           `gorm:"default:true"`
    FirstName      string
    LastName       string
    DateOfBirth    datatypes.Date
    Gender         Gender
    MaritalStatus  MaritalStatus
    OfficeDetail   []OfficeDetail `gorm:"foreignKey:ID"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    AddressDetail  []AddressDetail `gorm:"foreignKey:ID"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    CreatedAt      time.Time      
    UpdatedAt      time.Time       
   
}
