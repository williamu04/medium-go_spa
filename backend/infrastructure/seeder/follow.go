package seeder

import (
	"math/rand"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type FollowDataSeeder struct {
	count  int
	db     *gorm.DB
	logger *pkg.Logger
}

func NewFollowDataSeeder(count int, db *gorm.DB, logger *pkg.Logger) *FollowDataSeeder {
	return &FollowDataSeeder{
		count:  count,
		db:     db,
		logger: logger,
	}
}

func (s *FollowDataSeeder) Seed() error {
	var users []model.User

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		s.logger.Errorf("failed to fetch users: %v", err)
		return err
	}

	if len(users) < 2 {
		s.logger.Errorf("need at least 2 users to create follows, got %d", len(users))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	follows := make([]*model.Follow, 0, s.count)

	for range s.count {
		followerID := users[rand.Intn(len(users))].ID
		followingID := users[rand.Intn(len(users))].ID

		// Skip self-follows
		if followerID == followingID {
			continue
		}

		follows = append(follows, &model.Follow{
			FollowedByID: followerID,
			FollowingID:  followingID,
		})

	}

	if err := session.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(follows, 1000).Error; err != nil {
		return err
	}

	s.logger.Infof("✓ Seeded %d follows", s.count)
	return nil
}
