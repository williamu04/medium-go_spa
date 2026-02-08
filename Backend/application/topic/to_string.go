package topic_application

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
	outputs := make([]string, len(topics))

	for i, topic := range topics {
		outputs[i] = topic.Topic
	}

	return outputs, nil
}
