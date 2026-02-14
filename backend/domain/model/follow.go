package model

import "gorm.io/gorm"

type FollowModel struct {
	gorm.Model
	Following    *UserModel
	FollowingID  uint
	FollowedBy   *UserModel
	FollowedByID uint
}
