package topic_application

import (
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type TopicUseCase struct {
	Create      *CreateTopicUseCase
	RetrieveAll *RetrieveAllTopicUseCase
	Update      *UpdateTopicUseCase
	Delete      *DeleteTopicUseCase
	ToString    *ToStringUseCase
}

func NewTopicUseCase(repository repository.TopicRepository, sluger *pkg.Sluger) *TopicUseCase {
	return &TopicUseCase{
		Create:      NewCreateTopicUseCase(repository, sluger),
		RetrieveAll: NewRetrieveAllTopicUseCase(repository),
		Update:      NewUpdateTopicUseCase(repository, sluger),
		Delete:      NewDeleteTopicUseCase(repository),
		ToString:    NewToStringUseCase(repository),
	}
}
