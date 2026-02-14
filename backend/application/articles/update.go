package articleapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type UpdateArticleUseCase struct {
	repository repository.ArticleRepository
	topic      repository.TopicRepository
	sluger     *pkg.Sluger
}

func NewUpdateArticleUseCase(
	repository repository.ArticleRepository,
	topic repository.TopicRepository,
	sluger *pkg.Sluger,
) *UpdateArticleUseCase {
	return &UpdateArticleUseCase{
		repository: repository,
		topic:      topic,
		sluger:     sluger,
	}
}

type UpdateArticleInput struct {
	Title       string
	Description string
	Body        string
	Topic       []string
}

func (uc *UpdateArticleUseCase) Execute(ctx context.Context, input *UpdateArticleInput, ID uint) error {
	if ID == 0 {
		return nil
	}

	article, err := uc.repository.FindOneArticle(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if input.Title != "" {
		article.Title = input.Title
	}

	if input.Description != "" {
		article.Description = input.Description
	}

	if input.Body != "" {
		article.Body = input.Body
	}

	topics, err := uc.topic.FindAllTopics(ctx, input.Topic)
	if err != nil {
		return err
	}

	if topics == nil {
		return domain.ErrTopicNotFound
	}

	if input.Topic != nil {
		article.Topic = topics
	}

	article.Slug = uc.sluger.Slug(input.Title)

	return uc.repository.UpdateArticle(ctx, article, ID)
}
