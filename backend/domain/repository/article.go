package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type FeedArticles struct {
	ID            uint
	Title         string
	Description   string
	Thumbnail     *string
	AuthorAvatar  *string
	AuthorName    string
	CommentCount  int
	BookmarkCount int
}

type ArticleRepository interface {
	SaveOneArticle(ctx context.Context, article *model.Article) error
	FindOneArticle(ctx context.Context, filter map[string]any) (*model.Article, error)
	FeedArticles(ctx context.Context, limit int) ([]FeedArticles, error)
	FindAllArticles(ctx context.Context, filter map[string]any) ([]*model.Article, error)
	UpdateArticle(ctx context.Context, article *model.Article, id uint) error
	DeleteArticle(ctx context.Context, id uint) error
}
