package seeder

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

type UserDataSeeder struct {
	count  int
	hasher *pkg.Hasher
	db     *gorm.DB
	logger *pkg.Logger
}

func NewUserDataSeeder(count int, hasher *pkg.Hasher, db *gorm.DB, logger *pkg.Logger) *UserDataSeeder {
	return &UserDataSeeder{
		count:  count,
		hasher: hasher,
		db:     db,
		logger: logger,
	}
}

func (s *UserDataSeeder) Seed() error {
	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	for i := range s.count {
		// Hash default password
		passwordHash, err := s.hasher.Hash("password123")
		if err != nil {
			s.logger.Errorf("Failed to hash password: %v", err)
			continue
		}

		fakeName := faker.Username()
		fakeEmail := faker.Email()
		user := &model.UserModel{
			Username:     fmt.Sprintf("%d_%s", i, fakeName),
			Email:        fmt.Sprintf("%d_%s", i, fakeEmail),
			Bio:          faker.Sentence(),
			PasswordHash: passwordHash,
		}

		// Optional image (50% chance)
		if i%2 == 0 {
			img := fmt.Sprintf("https://i.pravatar.cc/150?img=%d", i)
			user.Image = &img
		}

		if err := session.Omit("Articles", "Comments", "Bookmarks", "Following", "FollowedBy").Create(user).Error; err != nil {
			s.logger.Errorf("Create user %d failed: %v", i, err)
			continue
		}
		//		s.logger.Infof("Created user %d: %s (ID=%d)", i, user.Username, user.ID)
	}

	s.logger.Infof("✓ Seeded %d users", s.count)
	return nil
}
