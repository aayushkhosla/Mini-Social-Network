package models

import (
    "gorm.io/gorm"
    "gorm.io/datatypes"
)

type AddressDetail struct {
    *gorm.Model
    ID           uint           `gorm:"primaryKey"`
    UserID       uint           `gorm:"index"`
    Address      string         `gorm:"size:255"`
    City         string
    State        string
    Country      string
    ContactNo1   string
    ContactNo2   string
    UpdatedAt    datatypes.Date
}
