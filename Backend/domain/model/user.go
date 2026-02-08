package model

import (
	"regexp"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username     string           `gorm:"column:username;uniqueIndex"`
	Email        string           `gorm:"column:email;uniqueIndex"`
	Bio          string           `gorm:"column:bio;size:1024"`
	Image        *string          `gorm:"column:image"`
	PasswordHash string           `gorm:"column:password;not null"`
	Articles     []*ArticleModel  `gorm:"foreignKey:AuthorID;references:ID"`
	Comments     []*CommentModel  `gorm:"foreignKey:AuthorID;references:ID"`
	Bookmarks    []*BookmarkModel `gorm:"ForeignKey:UserID"`
	Following    []*UserModel     `gorm:"many2many:user_follows;foreignKey:ID;joinForeignKey:follower_id;references:ID;joinReferences:following_id"`
	FollowedBy   []*UserModel     `gorm:"many2many:user_follows;foreignKey:ID;joinForeignKey:following_id;references:ID;joinReferences:follower_id"`
}

func (u *UserModel) IsEmailValid() bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email)
}

func (u *UserModel) IsPasswordValid(rawPassword string) bool {
	return len(rawPassword) >= 8
}
