package seeder

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/go-faker/faker/v4"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	var users []model.User
	var topics []model.Topic

	if err := s.db.Select("id").Find(&users).Error; err != nil {
		return err
	}
	if err := s.db.Select("id").Find(&topics).Error; err != nil {
		return err
	}

	if len(users) == 0 || len(topics) == 0 {
		return fmt.Errorf("insufficient data: users=%d topics=%d", len(users), len(topics))
	}

	session := s.db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	// =====================
	// 1️⃣ Generate Articles
	// =====================
	articles := make([]*model.Article, 0, s.count)

	for i := 0; i < s.count; i++ {
		title := cases.Title(language.Und).String(faker.Sentence())
		thumbnail := "https://placehold.co/300x200"

		articles = append(articles, &model.Article{
			Slug:        s.sluger.Slug(title),
			Title:       title,
			Description: faker.Paragraph(),
			Body:        strings.Repeat(faker.Paragraph()+" ", 3),
			Thumbnail:   &thumbnail,
			AuthorID:    users[rand.Intn(len(users))].ID,
		})
	}

	// =====================
	// 2️⃣ Bulk Insert Article
	// =====================
	if err := session.CreateInBatches(articles, 1000).Error; err != nil {
		return err
	}

	// =====================
	// 3️⃣ Generate Pivot Data
	// =====================
	var articleTopics []model.ArticleTopic

	for _, article := range articles {

		numTopics := rand.Intn(3) + 1 // 1-3 topics

		used := make(map[uint]bool)

		for range numTopics {
			topic := topics[rand.Intn(len(topics))]

			if used[topic.ID] {
				continue
			}

			used[topic.ID] = true

			articleTopics = append(articleTopics, model.ArticleTopic{
				ArticleID: article.ID,
				TopicID:   topic.ID,
			})
		}
	}

	// =====================
	// 4️⃣ Bulk Insert Pivot
	// =====================
	if err := session.CreateInBatches(articleTopics, 1000).Error; err != nil {
		return err
	}

	s.logger.Infof("✓ Seeded %d articles with topics", s.count)

	return nil
}
