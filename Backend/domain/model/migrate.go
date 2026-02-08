package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&FollowModel{})
	db.AutoMigrate(&ArticleModel{})
	db.AutoMigrate(&TopicModel{})
	db.AutoMigrate(&BookmarkModel{})
	db.AutoMigrate(&CommentModel{})
}
