package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body    string    `gorm:"size:2048"`
	Replies []Comment `gorm:"foreignKey:ParentID"`

	Parent    *Comment
	ParentID  *uint
	Article   *Article
	ArticleID uint
	Author    *User
	AuthorID  uint
}
