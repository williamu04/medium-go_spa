package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&TopicModel{})
	db.AutoMigrate(&ArticleModel{})
	db.AutoMigrate(&BookmarkModel{})
	db.AutoMigrate(&CommentModel{})
	db.AutoMigrate(&FollowModel{})
}
