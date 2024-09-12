package models

import (
	"time"

	// "gorm.io/datatypes"
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
    Email          string         `gorm:"unique;not null"`
    FirstName      string
    LastName       string
    DateOfBirth    time.Time   
    Gender         Gender
    MaritalStatus  MaritalStatus
    OfficeDetail   []OfficeDetail `gorm:"foreignKey:UserID"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    AddressDetail  []AddressDetail `gorm:"foreignKey:UserID"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  
}
