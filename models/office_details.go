package models

import (
    "gorm.io/gorm"
)

type OfficeDetail struct {
    gorm.Model
    ID            uint           `gorm:"primaryKey"`
    EmployeeCode  string         `gorm:"not null"`
    Address       string         `gorm:"size:255"`
    City          string
    State         string
    Country       string
    ContactNo     string
    OfficeEmail   string
    OfficeName    string
    UserID         uint           `gorm:"index"`
}




