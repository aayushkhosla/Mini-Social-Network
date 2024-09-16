package models

import (
	"gorm.io/gorm"
)

type AddressDetail struct {
    *gorm.Model
    ID           uint           `gorm:"primaryKey"`     
    Address      string         `gorm:"not null"`
    City         string         `gorm:"not null"`
    State        string         `gorm:"not null"`
    Country      string         `gorm:"not null"`
    ContactNo1   string         `gorm:"not null"`
    ContactNo2   string         
    UserID       uint            `gorm:"index"`    
}
