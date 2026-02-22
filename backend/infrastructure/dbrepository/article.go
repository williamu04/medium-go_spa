package dbrepository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewArticleDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DatabaseRepository) SaveOneArticle(ctx context.Context, article *model.Article) error {
	err := r.db.Save(article).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save article %s: %v", article.Title, err)
	}

	return err
}

func (r *DatabaseRepository) FindOneArticle(ctx context.Context, filter map[string]any) (*model.Article, error) {
	var article model.Article

	err := r.db.WithContext(ctx).Preload("Topic").Where(filter).First(&article).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if r.logger != nil {
				r.logger.Debugf("Article not found with filter: %v", filter)
			}
			return nil, nil
		}
		if r.logger != nil {
			r.logger.Errorf("Database error finding article with filter %v: %v", filter, err)
		}
		return nil, err
	}
	return &article, nil
}

func (r *DatabaseRepository) FindAllArticles(ctx context.Context, filter map[string]any) ([]*model.Article, error) {
	var articles []*model.Article
	return articles, r.db.Where(filter).Preload("Topic").Find(&articles).Error
}

func (r *DatabaseRepository) FeedArticles(
	ctx context.Context,
	limit int,
) ([]repository.FeedArticles, error) {
	var feeds []repository.FeedArticles

	query := `
	SELECT 
		a.id,
		a.title,
		a.description,
		a.thumbnail,
		u.name  AS author_name,
		u.image AS author_avatar,
		COALESCE(c.comment_count, 0)  AS comment_count,
		COALESCE(b.bookmark_count, 0) AS bookmark_count
	FROM articles a
	JOIN users u ON u.id = a.author_id

	LEFT JOIN (
		SELECT article_id, COUNT(*) AS comment_count
		FROM comments
		GROUP BY article_id
	) c ON c.article_id = a.id

	LEFT JOIN (
		SELECT article_id, COUNT(*) AS bookmark_count
		FROM bookmarks
		GROUP BY article_id
	) b ON b.article_id = a.id

	ORDER BY a.created_at DESC
	LIMIT $1
	`

	err := r.db.WithContext(ctx).
		Raw(query, limit).
		Scan(&feeds).Error

	return feeds, err
}

func (r *DatabaseRepository) UpdateArticle(ctx context.Context, article *model.Article, id uint) error {
	err := r.db.Model(article).Where("id = ?", id).Preload("Topic").Updates(article).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to update article ID %d: %v", id, err)
	}
	return err
}

func (r *DatabaseRepository) DeleteArticle(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Preload("Topic").Delete(&model.Article{}).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete article ID %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Article ID %d deleted", id)
	}
	return err
}
