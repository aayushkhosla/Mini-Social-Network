package models

import (
	"time"
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
    ID             uint            `gorm:"primaryKey"`
    Password       string          `gorm:"not null;json:-"`
    Email          string          `gorm:"unique;not null"`
    FirstName      string          `gorm:"not null"`
    LastName       string       
    DateOfBirth    time.Time       `gorm:"type:date;not null"`
    Gender         Gender          `gorm:"not null"`
    MaritalStatus  MaritalStatus   `gorm:"not null"`
    OfficeDetail   []OfficeDetail  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" `
    AddressDetail  []AddressDetail `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" `  
}
