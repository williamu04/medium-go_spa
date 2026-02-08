package follow_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveFollowUseCase struct {
	repository repository.FollowRepository
}

func NewRetrieveFollowUseCase(repository repository.FollowRepository) *RetrieveFollowUseCase {
	return &RetrieveFollowUseCase{repository: repository}
}

type RetrieveFollowOutput struct {
	ID           uint
	FollowingID  uint
	FollowedByID uint
}

func (uc *RetrieveFollowUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveFollowOutput, error) {
	Follow, err := uc.repository.FindOneFollow(ctx, filter)

	if err != nil {
		return nil, err
	}

	if Follow == nil {
		return nil, nil
	}

	return &RetrieveFollowOutput{
		ID:           Follow.ID,
		FollowingID:  Follow.FollowingID,
		FollowedByID: Follow.FollowedByID,
	}, nil
}
