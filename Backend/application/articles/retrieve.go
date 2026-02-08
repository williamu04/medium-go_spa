package article_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveArticleUseCase struct {
	repository repository.ArticleRepository
	topic      repository.TopicRepository
}

func NewRetrieveArticleUseCase(repository repository.ArticleRepository, topic repository.TopicRepository) *RetrieveArticleUseCase {
	return &RetrieveArticleUseCase{repository: repository, topic: topic}
}

type RetrieveArticleOutput struct {
	ID          uint
	Slug        string
	Title       string
	Description string
	Body        string
	AuthorID    uint
	Topic       []string
}

func (uc *RetrieveArticleUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveArticleOutput, error) {
	article, err := uc.repository.FindOneArticle(ctx, filter)
	if err != nil {
		return nil, err
	}

	if article == nil {
		return nil, nil
	}

	topics, err := uc.topic.SetToString(ctx, article.Topic)

	return &RetrieveArticleOutput{
		ID:          article.ID,
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		AuthorID:    article.AuthorID,
		Topic:       topics,
	}, nil
}
