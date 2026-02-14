package seeder

import (
	"math/rand"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
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
	var users []model.UserModel

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		s.logger.Errorf("failed to fetch users: %v", err)
		return err
	}

	if len(users) < 2 {
		s.logger.Errorf("need at least 2 users to create follows, got %d", len(users))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	created := 0
	attempts := 0
	maxAttempts := s.count * 3

	for created < s.count && attempts < maxAttempts {
		attempts++
		followerID := users[rand.Intn(len(users))].ID
		followingID := users[rand.Intn(len(users))].ID

		// Skip self-follows
		if followerID == followingID {
			continue
		}

		follow := &model.FollowModel{
			FollowedByID: followerID,
			FollowingID:  followingID,
		}

		if err := session.Omit("Following", "FollowedBy").Create(follow).Error; err != nil {
			// Skip duplicates
			continue
		}
		created++
		//		s.logger.Infof("Created follow %d: User %d → User %d (ID=%d)", created, followerID, followingID, follow.ID)
	}

	s.logger.Infof("✓ Seeded %d follows", created)
	return nil
}
