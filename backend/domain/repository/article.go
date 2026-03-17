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

type DetailedArticle struct {
	ID           uint
	Title        string
	Description  string
	Body         string
	Thumbnail    *string
	AuthorAvatar *string
	AuthorName   string
	Comments     []*model.Comment
}

type ArticleRepository interface {
	SaveOneArticle(ctx context.Context, article *model.Article) error
	FindOneArticle(ctx context.Context, filter map[string]any) (*model.Article, error)
	DetailedArticle(ctx context.Context, filter map[string]any) (DetailedArticle, error)
	FeedArticles(ctx context.Context, limit int) ([]FeedArticles, error)
	FindAllArticles(ctx context.Context, filter map[string]any) ([]*model.Article, error)
	UpdateArticle(ctx context.Context, article *model.Article, id uint) error
	DeleteArticle(ctx context.Context, id uint) error
}
