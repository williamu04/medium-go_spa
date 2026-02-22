package topicapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllTopicUseCase struct {
	repository repository.TopicRepository
}

func NewRetrieveAllTopicUseCase(repository repository.TopicRepository) *RetrieveAllTopicUseCase {
	return &RetrieveAllTopicUseCase{repository: repository}
}

type RetrieveTopicOutput struct {
	ID    uint
	Slug  string
	Topic string
}

type RetrieveAllTopicsOutput struct {
	Topics []RetrieveTopicOutput
}

func (uc *RetrieveAllTopicUseCase) Execute(ctx context.Context, slugs []string) (*RetrieveAllTopicsOutput, error) {
	topics, err := uc.repository.FindAllTopics(ctx, slugs)
	if err != nil {
		return nil, err
	}

	if topics == nil {
		return nil, nil
	}

	return &RetrieveAllTopicsOutput{
		Topics: mapTopicsToOutputs(topics),
	}, nil
}

func mapTopicsToOutputs(topics []*model.Topic) []RetrieveTopicOutput {
	outputs := make([]RetrieveTopicOutput, len(topics))

	for i, topic := range topics {
		outputs[i] = RetrieveTopicOutput{
			ID:    topic.ID,
			Slug:  topic.Slug,
			Topic: topic.Topic,
		}
	}
	return outputs
}
