package models

import (
	"gorm.io/gorm"
)

type Follow struct {
    gorm.Model
    UserID        uint           `gorm:"index"`
    FollowedUserID uint          `gorm:"index"`
    Active        bool           `gorm:"default:true"`
}
