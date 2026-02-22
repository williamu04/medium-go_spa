package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Following    *User
	FollowingID  uint `gorm:"UniqueIndex"`
	FollowedBy   *User
	FollowedByID uint `gorm:"UniqueIndex"`
}
