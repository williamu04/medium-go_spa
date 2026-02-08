package model

import "gorm.io/gorm"

type BookmarkModel struct {
	gorm.Model
	Article *ArticleModel
	ArticleID uint
	User *UserModel
	UserID uint
}