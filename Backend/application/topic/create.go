package topic_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type CreateTopicUseCase struct {
	repository repository.TopicRepository
	sluger     *pkg.Sluger
}

func NewCreateTopicUseCase(
	repository repository.TopicRepository,
	sluger *pkg.Sluger,
) *CreateTopicUseCase {
	return &CreateTopicUseCase{
		repository: repository,
		sluger:     sluger,
	}
}

type CreateTopicInput struct {
	Topic string
}

type CreateTopicOutput struct {
	ID    uint
	Slug  string
	Topic string
}

func (uc *CreateTopicUseCase) Execute(ctx context.Context, input *CreateTopicInput) (*CreateTopicOutput, error) {
	topic := &model.TopicModel{
		Topic: input.Topic,
	}

	topic.Slug = uc.sluger.Slug(input.Topic)

	if err := uc.repository.SaveOneTopic(ctx, topic); err != nil {
		return nil, err
	}

	return &CreateTopicOutput{
		ID:    topic.ID,
		Slug:  topic.Slug,
		Topic: topic.Topic,
	}, nil
}
