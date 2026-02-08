package article_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteArticleUseCase struct {
	repository repository.ArticleRepository
}

func NewDeleteArticleUseCase(repository repository.ArticleRepository) *DeleteArticleUseCase {
	return &DeleteArticleUseCase{repository: repository}
}

func (uc *DeleteArticleUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return nil
	}

	article, err := uc.repository.FindOneArticle(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if article == nil {
		return nil
	}

	return uc.repository.DeleteArticle(ctx, ID)
}
