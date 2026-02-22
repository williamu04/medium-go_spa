package model

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Slug     string `gorm:"uniqueIndex"`
	Topic    string
	Articles []*Article `gorm:"many2many:article_topics;"`
}
