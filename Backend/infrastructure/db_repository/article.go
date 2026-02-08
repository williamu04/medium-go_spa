package db_repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewArticleDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DatabaseRepository) SaveOneArticle(ctx context.Context, article *model.ArticleModel) error {
	err := r.db.Save(article).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save article %s: %v", article.Title, err)
	}

	return err
}

func (r *DatabaseRepository) FindOneArticle(ctx context.Context, filter map[string]any) (*model.ArticleModel, error) {
	var article model.ArticleModel

	err := r.db.WithContext(ctx).Where(filter).First(&article).Error

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

func (r *DatabaseRepository) FindAllArticles(ctx context.Context, filter map[string]any) ([]*model.ArticleModel, error) {
	var articles []*model.ArticleModel
	return articles, r.db.Where(filter).Find(&articles).Error
}

func (r *DatabaseRepository) UpdateArticle(ctx context.Context, article *model.ArticleModel, id uint) error {
	err := r.db.Model(article).Where("id = ?", id).Updates(article).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to update article ID %d: %v", id, err)
	}
	return err
}

func (r *DatabaseRepository) DeleteArticle(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.ArticleModel{}).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete article ID %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Article ID %d deleted", id)
	}
	return err
}
