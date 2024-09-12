package models

import (
	"gorm.io/gorm"
)

type AddressDetail struct {
    *gorm.Model
    ID           uint           `gorm:"primaryKey"`     
    Address      string         `gorm:"size:255"`
    City         string
    State        string
    Country      string
    ContactNo1   string
    ContactNo2   string
    UserID       uint            `gorm:"index"`    
}
