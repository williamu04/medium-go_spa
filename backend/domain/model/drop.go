package model

import "gorm.io/gorm"

func DropAll(db *gorm.DB) {
	db.Migrator().DropTable(&UserModel{})
	db.Migrator().DropTable(&TopicModel{})
	db.Migrator().DropTable(&ArticleModel{})
	db.Migrator().DropTable(&CommentModel{})
	db.Migrator().DropTable(&BookmarkModel{})
	db.Migrator().DropTable(&FollowModel{})
}

func DropAllQuery(db *gorm.DB) {
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}
