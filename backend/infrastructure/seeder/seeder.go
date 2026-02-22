package seeder

import (
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

type SeedHandler struct {
	userSeeder     *UserDataSeeder
	topicSeeder    *TopicDataSeeder
	articleSeeder  *ArticleDataSeeder
	commentSeeder  *CommentDataSeeder
	bookmarkSeeder *BookmarkDataSeeder
	followSeeder   *FollowDataSeeder
	logger         *pkg.Logger
}

func NewSeedHandler(db *gorm.DB, logger *pkg.Logger, hasher *pkg.Hasher, sluger *pkg.Sluger, randomer *pkg.Randomer) *SeedHandler {
	return &SeedHandler{
		userSeeder:     NewUserDataSeeder(1000, hasher, sluger, db, logger),
		topicSeeder:    NewTopicDataSeeder(200, sluger, db, logger, randomer),
		articleSeeder:  NewArticleDataSeeder(20000, sluger, db, logger),
		commentSeeder:  NewCommentDataSeeder(100000, db, logger),
		bookmarkSeeder: NewBookmarkDataSeeder(400000, db, logger),
		followSeeder:   NewFollowDataSeeder(100000, db, logger),
		logger:         logger,
	}
}

func (s *SeedHandler) SeedAll() error {
	// URUTAN PENTING: Users & Topics dulu (independent)
	if err := s.userSeeder.Seed(); err != nil {
		s.logger.Errorf("user seeding failed: %v", err)
		return err
	}

	if err := s.followSeeder.Seed(); err != nil {
		s.logger.Errorf("follow seeding failed: %v", err)
		return err
	}

	if err := s.topicSeeder.Seed(); err != nil {
		s.logger.Errorf("topic seeding failed: %v", err)
		return err
	}

	// Articles butuh Users & Topics
	if err := s.articleSeeder.Seed(); err != nil {
		s.logger.Errorf("article seeding failed: %v", err)
		return err
	}

	// Bookmarks & Follows bisa parallel (keduanya butuh Users & Articles)
	if err := s.bookmarkSeeder.Seed(); err != nil {
		s.logger.Errorf("bookmark seeding failed: %v", err)
		return err
	}

	// Comments butuh Users & Articles
	if err := s.commentSeeder.Seed(); err != nil {
		s.logger.Errorf("comment seeding failed: %v", err)
		return err
	}
	s.logger.Info("🎉 All data seeded successfully!")
	return nil
}
