package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Slug         string `gorm:"uniqueIndex"`
	Title        string
	Description  string `gorm:"size:2048"`
	Body         string `gorm:"size:2048"`
	Thumbnail    *string
	Author       *User
	AuthorID     uint
	Topics       []*Topic    `gorm:"many2many:article_topics;"`
	Comments     []*Comment  `gorm:"ForeignKey:ArticleID"`
	BookmarkedBy []*Bookmark `gorm:"ForeignKey:ArticleID"`
}

type ArticleTopic struct {
	ArticleID uint
	TopicID   uint
}

func (a *Article) IsTopicValid(topic []string) bool {
	return len(topic) > 0 && len(topic) <= 5
}
