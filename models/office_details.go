package models

import (
    "gorm.io/gorm"
)

type OfficeDetail struct {
    gorm.Model
    ID            uint              `gorm:"primaryKey"`
    EmployeeCode  string            `gorm:"not null,uniqueIndex:idx_name"`
    Address       string            `gorm:"not null"`
    City          string            `gorm:"not null"`
    State         string            `gorm:"not null"`
    Country       string            `gorm:"not null"`
    ContactNo     string            `gorm:"not null"`
    OfficeEmail   string            `gorm:"not null,uniqueIndex:idx_name"`
    OfficeName    string            `gorm:"not null"`
    UserID         uint             `gorm:"index"`
}




