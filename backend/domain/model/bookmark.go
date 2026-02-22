package model

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	Article   *Article
	ArticleID uint `gorm:"UniqueIndex"`
	User      *User
	UserID    uint `gorm:"UniqueIndex"`
}
