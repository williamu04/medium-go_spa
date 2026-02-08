package model

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	Body      string `gorm:"size:2048"`
	Article   *ArticleModel
	ArticleID uint
	Author    *UserModel
	AuthorID  uint
}
