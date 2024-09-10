package models

import (
    "gorm.io/gorm"
    "gorm.io/datatypes"
)

// Follow represents a relationship where one user follows another
type Follow struct {
    gorm.Model
    UserID        uint           `gorm:"index"`
    FollowedUserID uint          `gorm:"index"`
    Active        bool           `gorm:"default:true"`
    CreatedAt     datatypes.Date
}
