package model

import "gorm.io/gorm"

func DropAll(db *gorm.DB) {
	db.Migrator().DropTable(&User{})
	db.Migrator().DropTable(&Topic{})
	db.Migrator().DropTable(&Article{})
	db.Migrator().DropTable(&Comment{})
	db.Migrator().DropTable(&Bookmark{})
	db.Migrator().DropTable(&Follow{})
}

func DropAllQuery(db *gorm.DB) {
	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}
