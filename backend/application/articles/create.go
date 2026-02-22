package articleapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type CreateArticleUseCase struct {
	repository repository.ArticleRepository
	topic      repository.TopicRepository
	sluger     *pkg.Sluger
}

func NewCreateArticleUseCase(
	repository repository.ArticleRepository,
	topic repository.TopicRepository,
	sluger *pkg.Sluger,
) *CreateArticleUseCase {
	return &CreateArticleUseCase{
		repository: repository,
		topic:      topic,
		sluger:     sluger,
	}
}

type CreateArticleInput struct {
	Title       string
	Description string
	Body        string
	AuthorID    uint
	Topic       []string
}

type CreateArticleOutput struct {
	ID          uint
	Slug        string
	Title       string
	Description string
	Body        string
	AuthorID    uint
	Topic       []string
}

func (uc *CreateArticleUseCase) Execute(ctx context.Context, input *CreateArticleInput) (*CreateArticleOutput, error) {
	topics, err := uc.topic.FindAllTopics(ctx, input.Topic)
	if err != nil {
		return nil, err
	}

	if topics == nil {
		return nil, domain.ErrTopicNotFound
	}

	article := &model.Article{
		Title:       input.Title,
		Description: input.Description,
		Body:        input.Body,
		AuthorID:    input.AuthorID,
	}

	article.Slug = uc.sluger.Slug(input.Title)

	if err := uc.repository.SaveOneArticle(ctx, article); err != nil {
		return nil, err
	}

	topicStr, err := uc.topic.SetToString(ctx, topics)
	if err != nil {
		return nil, err
	}

	return &CreateArticleOutput{
		ID:          article.ID,
		Slug:        article.Slug,
		Title:       article.Title,
		Description: article.Description,
		Body:        article.Body,
		AuthorID:    article.AuthorID,
		Topic:       topicStr,
	}, nil
}
