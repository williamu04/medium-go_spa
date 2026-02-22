package seeder

import (
	"fmt"
	"math/rand"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type BookmarkDataSeeder struct {
	count  int
	db     *gorm.DB
	logger *pkg.Logger
}

func NewBookmarkDataSeeder(count int, db *gorm.DB, logger *pkg.Logger) *BookmarkDataSeeder {
	return &BookmarkDataSeeder{
		count:  count,
		db:     db,
		logger: logger,
	}
}

func (s *BookmarkDataSeeder) Seed() error {
	var users []model.User
	var articles []model.Article

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		return err
	}
	if err := s.db.Select("id").Find(&articles).Error; err != nil {
		return err
	}

	if len(users) == 0 || len(articles) == 0 {
		return fmt.Errorf("insufficient data: users=%d articles=%d",
			len(users), len(articles))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	bookmarks := make([]*model.Bookmark, 0, s.count)

	for i := 0; i < s.count; i++ {
		bookmarks = append(bookmarks, &model.Bookmark{
			UserID:    users[rand.Intn(len(users))].ID,
			ArticleID: articles[rand.Intn(len(articles))].ID,
		})
	}

	if err := session.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(bookmarks, 1000).Error; err != nil {
		return err
	}

	s.logger.Infof("✓ Seeded up to %d bookmarks (duplicates ignored)", s.count)

	return nil
}
