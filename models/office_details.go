package models

import (
    "gorm.io/gorm"
)

type OfficeDetail struct {
    gorm.Model
    ID            uint           `gorm:"primaryKey"`
    UserID        uint           `gorm:"index"`
    EmployeeCode  string         `gorm:"not null"`
    Address       string         `gorm:"size:255"`
    City          string
    State         string
    Country       string
    ContactNo     string
    Email         string
    OfficeName    string
}




