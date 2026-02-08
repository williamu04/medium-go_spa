package model

import "gorm.io/gorm"

type TopicModel struct {
	gorm.Model
	Slug         string `gorm:"uniqueIndex"`
	Topic        string
	ArticleModel []*ArticleModel `gorm:"many2many:article_topic;"`
}
