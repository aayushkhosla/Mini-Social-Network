package models

import (
	"gorm.io/gorm"
)

type Follow struct {
    gorm.Model
    UserID        uint           `gorm:"index not null"`
    FollowedUserID uint          `gorm:"index not null"`
    Active        bool           `gorm:"default:true"`
    User          User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
    FollowedUser  User           `gorm:"foreignKey:FollowedUserID;constraint:OnDelete:CASCADE;"`
}
