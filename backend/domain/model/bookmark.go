package model

import "gorm.io/gorm"

type Bookmark struct {
	gorm.Model
	Article   *Article
	ArticleID uint `gorm:"uniqueIndex:idx_user_article"`
	User      *User
	UserID    uint `gorm:"uniqueIndex:idx_user_article"`
}
