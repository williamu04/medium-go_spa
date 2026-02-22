package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Topic{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Bookmark{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Follow{})
}
