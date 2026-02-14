package topicapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type UpdateTopicUseCase struct {
	repository repository.TopicRepository
	sluger     *pkg.Sluger
}

func NewUpdateTopicUseCase(
	repository repository.TopicRepository,
	sluger *pkg.Sluger,
) *UpdateTopicUseCase {
	return &UpdateTopicUseCase{
		repository: repository,
		sluger:     sluger,
	}
}

type UpdateTopicInput struct {
	Topic string
}

func (uc *UpdateTopicUseCase) Execute(ctx context.Context, input *UpdateTopicInput, ID uint) error {
	if ID == 0 {
		return nil
	}

	topic, err := uc.repository.FindOneTopic(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if input.Topic != "" {
		topic.Topic = input.Topic
	}

	topic.Slug = uc.sluger.Slug(input.Topic)

	return uc.repository.UpdateTopic(ctx, topic, ID)
}
