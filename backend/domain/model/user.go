package model

import (
	"regexp"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Username     string `gorm:"uniqueIndex"`
	Slug         string `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
	Bio          string `gorm:"size:1024"`
	Image        *string
	PasswordHash string      `gorm:"not null"`
	Articles     []*Article  `gorm:"foreignKey:AuthorID"`
	Comments     []*Comment  `gorm:"foreignKey:AuthorID"`
	Bookmarks    []*Bookmark `gorm:"ForeignKey:UserID"`
	Following    []*User     `gorm:"many2many:user_follows;foreignKey:ID;joinForeignKey:follower_id;references:ID;joinReferences:following_id"`
	FollowedBy   []*User     `gorm:"many2many:user_follows;foreignKey:ID;joinForeignKey:following_id;references:ID;joinReferences:follower_id"`
}

func (u *User) IsEmailValid() bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(u.Email)
}

func (u *User) IsPasswordValid(rawPassword string) bool {
	return len(rawPassword) >= 8
}
