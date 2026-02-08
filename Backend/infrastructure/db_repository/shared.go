package db_repository

import (
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

type DatabaseRepository struct {
	db     *gorm.DB
	logger *pkg.Logger
}
