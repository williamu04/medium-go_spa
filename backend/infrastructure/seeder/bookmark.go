package seeder

import (
	"math/rand"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
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
	var users []model.UserModel
	var articles []model.ArticleModel

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		s.logger.Errorf("failed to fetch users: %v", err)
		return err
	}
	if err := s.db.Select("id").Find(&articles).Error; err != nil {
		s.logger.Errorf("failed to fetch articles: %v", err)
		return err
	}

	if len(users) == 0 || len(articles) == 0 {
		s.logger.Errorf("insufficient data: users=%d, articles=%d", len(users), len(articles))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	created := 0
	attempts := 0
	maxAttempts := s.count * 3 // Avoid infinite loop on duplicates

	for created < s.count && attempts < maxAttempts {
		attempts++
		bookmark := &model.BookmarkModel{
			UserID:    users[rand.Intn(len(users))].ID,
			ArticleID: articles[rand.Intn(len(articles))].ID,
		}

		if err := session.Omit("User", "Article").Create(bookmark).Error; err != nil {
			// Skip duplicates silently
			continue
		}
		created++
		//		s.logger.Infof("Created bookmark %d (ID=%d)", created, bookmark.ID)
	}

	s.logger.Infof("✓ Seeded %d bookmarks", created)
	return nil
}
