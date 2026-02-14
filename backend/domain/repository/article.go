package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type ArticleRepository interface {
	SaveOneArticle(ctx context.Context, article *model.ArticleModel) error
	FindOneArticle(ctx context.Context, filter map[string]any) (*model.ArticleModel, error)
	FindAllArticles(ctx context.Context, filter map[string]any) ([]*model.ArticleModel, error)
	UpdateArticle(ctx context.Context, article *model.ArticleModel, id uint) error
	DeleteArticle(ctx context.Context, id uint) error
}
