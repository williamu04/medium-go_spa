package topicapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteTopicUseCase struct {
	repository repository.TopicRepository
}

func NewDeleteTopicUseCase(repository repository.TopicRepository) *DeleteTopicUseCase {
	return &DeleteTopicUseCase{repository: repository}
}

func (uc *DeleteTopicUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return nil
	}

	topic, err := uc.repository.FindOneTopic(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if topic == nil {
		return nil
	}

	return uc.repository.DeleteTopic(ctx, ID)
}
