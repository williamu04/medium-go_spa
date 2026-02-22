package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body      string `gorm:"size:2048"`
	Article   *Article
	ArticleID uint
	Author    *User
	AuthorID  uint
}
