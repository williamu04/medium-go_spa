package topicapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type ToStringUseCase struct {
	repository repository.TopicRepository
}

func NewToStringUseCase(repository repository.TopicRepository) *ToStringUseCase {
	return &ToStringUseCase{repository: repository}
}

func (uc *ToStringUseCase) Execute(ctx context.Context, topics []*model.TopicModel) ([]string, error) {
	topicStr, err := uc.repository.SetToString(ctx, topics)
	if err != nil {
		return nil, err
	}

	return topicStr, nil
}
