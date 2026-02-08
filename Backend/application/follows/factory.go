package follow_application

import "github.com/williamu04/medium-clone/domain/repository"

type FollowUseCase struct {
	Create      *CreateFollowUseCase
	Retrieve    *RetrieveFollowUseCase
	RetrieveAll *RetrieveAllFollowUseCase
	Delete      *DeleteFollowUseCase
}

func NewFollowUseCase(repository repository.FollowRepository) *FollowUseCase {
	return &FollowUseCase{
		Create:      NewCreateFollowUseCase(repository),
		Retrieve:    NewRetrieveFollowUseCase(repository),
		RetrieveAll: NewRetrieveAllFollowUseCase(repository),
		Delete:      NewDeleteFollowUseCase(repository),
	}
}
