package seeder

import (
	"math/rand"
	"strings"

	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type ArticleDataSeeder struct {
	count  int
	sluger *pkg.Sluger
	db     *gorm.DB
	logger *pkg.Logger
}

func NewArticleDataSeeder(count int, sluger *pkg.Sluger, db *gorm.DB, logger *pkg.Logger) *ArticleDataSeeder {
	return &ArticleDataSeeder{
		count:  count,
		sluger: sluger,
		db:     db,
		logger: logger,
	}
}

func (s *ArticleDataSeeder) Seed() error {
	// Get all users and topics
	var users []model.UserModel
	var topics []model.TopicModel

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		s.logger.Errorf("failed to fetch users: %v", err)
		return err
	}
	if err := s.db.Select("id, slug").Find(&topics).Error; err != nil {
		s.logger.Errorf("failed to fetch topics: %v", err)
		return err
	}

	if len(users) == 0 || len(topics) == 0 {
		s.logger.Errorf("insufficient data: users=%d, topics=%d", len(users), len(topics))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	for i := range s.count {
		title := cases.Title(language.Und).String(faker.Sentence())
		article := &model.ArticleModel{
			Slug:        s.sluger.Slug(title),
			Title:       title,
			Description: faker.Paragraph(),
			Body:        strings.Repeat(faker.Paragraph()+" ", 3),
			AuthorID:    users[rand.Intn(len(users))].ID,
		}

		// Create article without associations first
		if err := session.Omit("Author", "Topic", "Comments", "BookmarkedBy").Create(article).Error; err != nil {
			s.logger.Errorf("Create article %d failed: %v", i, err)
			continue
		}

		// Assign 1-3 random topics
		numTopics := rand.Intn(5) + 1
		selectedTopics := make([]*model.TopicModel, 0, numTopics)
		for range numTopics {
			selectedTopics = append(selectedTopics, &topics[rand.Intn(len(topics))])
		}

		if err := s.db.Model(article).Association("Topic").Append(selectedTopics); err != nil {
			s.logger.Warnf("Failed to append topics to article %d: %v", article.ID, err)
		}

		// s.logger.Infof("Created article %d: %s (ID=%d)", i, article.Title, article.ID)
	}

	s.logger.Infof("✓ Seeded %d articles", s.count)
	return nil
}
