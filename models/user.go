package models

import (
    "gorm.io/gorm"
    "gorm.io/datatypes"
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
    CreatedAt      datatypes.Date
    UpdatedAt      datatypes.Date
    OfficeDetails  []OfficeDetail `gorm:"foreignKey:UserID"`
    AddressDetails []AddressDetail `gorm:"foreignKey:UserID"`
   
}
