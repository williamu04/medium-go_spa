package seeder

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserDataSeeder struct {
	count  int
	hasher *pkg.Hasher
	sluger *pkg.Sluger
	db     *gorm.DB
	logger *pkg.Logger
}

func NewUserDataSeeder(count int, hasher *pkg.Hasher, sluger *pkg.Sluger, db *gorm.DB, logger *pkg.Logger) *UserDataSeeder {
	return &UserDataSeeder{
		count:  count,
		hasher: hasher,
		sluger: sluger,
		db:     db,
		logger: logger,
	}
}

func (s *UserDataSeeder) Seed() error {
	passwordHash, err := s.hasher.Hash("password123")
	if err != nil {
		s.logger.Errorf("Failed to hash password: %v", err)
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	users := make([]*model.User, 0, s.count)

	for i := range s.count {
		// Hash default password
		fakeName := faker.Name()
		slugName := s.sluger.Slug(fakeName)
		fakeUsername := faker.Username()
		fakeEmail := faker.Email()
		img := fmt.Sprintf("https://i.pravatar.cc/150?img=%d", i)

		users = append(users, &model.User{
			Name:         fmt.Sprintf("%d_%s", i, fakeName),
			Slug:         fmt.Sprintf("%d_%s", i, slugName),
			Username:     fmt.Sprintf("%d_%s", i, fakeUsername),
			Email:        fmt.Sprintf("%d_%s", i, fakeEmail),
			Bio:          faker.Sentence(),
			Image:        &img,
			PasswordHash: passwordHash,
		})
	}

	if err := session.CreateInBatches(users, 1000).Error; err != nil {
		return err
	}

	s.logger.Infof("✓ Seeded %d users", s.count)
	return nil
}
