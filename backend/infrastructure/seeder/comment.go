package seeder

import (
	"math/rand"

	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

type CommentDataSeeder struct {
	count  int
	db     *gorm.DB
	logger *pkg.Logger
}

func NewCommentDataSeeder(count int, db *gorm.DB, logger *pkg.Logger) *CommentDataSeeder {
	return &CommentDataSeeder{
		count:  count,
		db:     db,
		logger: logger,
	}
}

func (s *CommentDataSeeder) Seed() error {
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

	for i := range s.count {
		comment := &model.CommentModel{
			Body:      faker.Paragraph(),
			ArticleID: articles[rand.Intn(len(articles))].ID,
			AuthorID:  users[rand.Intn(len(users))].ID,
		}

		if err := session.Omit("Article", "Author").Create(comment).Error; err != nil {
			s.logger.Errorf("Create comment %d failed: %v", i, err)
			continue
		}
		// s.logger.Infof("Created comment %d (ID=%d)", i, comment.ID)
	}

	s.logger.Infof("✓ Seeded %d comments", s.count)
	return nil
}
