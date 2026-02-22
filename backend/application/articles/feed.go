package articleapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type FeedArticleUseCase struct {
	article  repository.ArticleRepository
	user     repository.UserRepository
	comment  repository.CommentRepository
	bookmark repository.BookmarkRepository
}

func NewFeedArticleUseCase(
	article repository.ArticleRepository,
	user repository.UserRepository,
	comment repository.CommentRepository,
	bookmark repository.BookmarkRepository,
) *FeedArticleUseCase {
	return &FeedArticleUseCase{
		article:  article,
		user:     user,
		comment:  comment,
		bookmark: bookmark,
	}
}

type FeedArticlesOutput struct {
	Articles []repository.FeedArticles
}

func (uc *FeedArticleUseCase) Execute(
	ctx context.Context,
	limit int,
) (*FeedArticlesOutput, error) {
	feeds, err := uc.article.FeedArticles(ctx, limit)
	if err != nil {
		return nil, err
	}

	return &FeedArticlesOutput{
		Articles: feeds,
	}, nil
}
