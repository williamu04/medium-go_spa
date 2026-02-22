package seeder

import (
	"math/rand"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TopicDataSeeder struct {
	count    int
	sluger   *pkg.Sluger
	db       *gorm.DB
	logger   *pkg.Logger
	randomer *pkg.Randomer
}

func NewTopicDataSeeder(count int, sluger *pkg.Sluger, db *gorm.DB, logger *pkg.Logger, randomer *pkg.Randomer) *TopicDataSeeder {
	return &TopicDataSeeder{
		count:    count,
		sluger:   sluger,
		db:       db,
		logger:   logger,
		randomer: randomer,
	}
}

func (s *TopicDataSeeder) Seed() error {
	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	for i := range s.count {
		fakeTopic := cases.Title(language.Und).String(s.randomer.RandomWord(rand.Intn(3) + 5))
		topic := &model.Topic{
			// Topic: topics[i],
			Topic: fakeTopic,
			Slug:  s.sluger.Slug(fakeTopic),
		}

		if err := session.Omit("Article").Create(topic).Error; err != nil {
			s.logger.Errorf("Create topic %d failed: %v", i, err)
			continue
		}
		//		s.logger.Infof("Created topic %d: %s (ID=%d)", i, topic.Topic, topic.ID)
	}

	s.logger.Infof("✓ Seeded %d topics", s.count)
	return nil
}
