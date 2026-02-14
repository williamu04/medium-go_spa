package model

import "gorm.io/gorm"

type ArticleModel struct {
	gorm.Model
	Slug         string `gorm:"uniqueIndex"`
	Title        string
	Description  string `gorm:"size:2048"`
	Body         string `gorm:"size:2048"`
	Author       *UserModel
	AuthorID     uint
	Topic        []*TopicModel    `gorm:"many2many:article_topics;"`
	Comments     []*CommentModel  `gorm:"ForeignKey:ArticleID;references:ID"`
	BookmarkedBy []*BookmarkModel `gorm:"ForeignKey:ArticleID"`
}

func (a *ArticleModel) IsTopicValid(topic []string) bool {
	return len(topic) > 0 && len(topic) <= 5
}
