package seeder

import (
	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type TopicDataSeeder struct {
	count  int
	sluger *pkg.Sluger
	db     *gorm.DB
	logger *pkg.Logger
}

func NewTopicDataSeeder(count int, sluger *pkg.Sluger, db *gorm.DB, logger *pkg.Logger) *TopicDataSeeder {
	return &TopicDataSeeder{
		count:  count,
		sluger: sluger,
		db:     db,
		logger: logger,
	}
}

func (s *TopicDataSeeder) Seed() error {
	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	topics := []string{
		"Technology", "Science", "Health", "Business", "Politics",
		"Entertainment", "Sports", "Travel", "Food", "Fashion",
		"Education", "Environment", "Finance", "Gaming", "Music",
		"Art", "Photography", "Programming", "AI", "Blockchain",
	}

	for i := range s.count {
		fakeTopic := cases.Title(language.Und).String(faker.Word())
		topic := &model.TopicModel{
			// Topic: topics[i],
			Topic: fakeTopic,
			Slug:  s.sluger.Slug(fakeTopic),
		}

		if err := session.Omit("ArticleModel").Create(topic).Error; err != nil {
			s.logger.Errorf("Create topic %d failed: %v", i, err)
			continue
		}
		//		s.logger.Infof("Created topic %d: %s (ID=%d)", i, topic.Topic, topic.ID)
	}

	s.logger.Infof("✓ Seeded %d topics", min(s.count, len(topics)))
	return nil
}
