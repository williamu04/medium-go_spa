package followapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type CreateFollowUseCase struct {
	repository repository.FollowRepository
}

func NewCreateFollowUseCase(repository repository.FollowRepository) *CreateFollowUseCase {
	return &CreateFollowUseCase{repository: repository}
}

type CreateFollowInput struct {
	FollowingID  uint
	FollowedByID uint
}

type CreateFollowOutput struct {
	ID           uint
	FollowingID  uint
	FollowedByID uint
}

func (uc *CreateFollowUseCase) Execute(ctx context.Context, input *CreateFollowInput) (*CreateFollowOutput, error) {
	follow := &model.Follow{
		FollowingID:  input.FollowingID,
		FollowedByID: input.FollowedByID,
	}

	if err := uc.repository.SaveOneFollow(ctx, follow); err != nil {
		return nil, err
	}

	return &CreateFollowOutput{
		ID:           follow.ID,
		FollowingID:  follow.FollowingID,
		FollowedByID: follow.FollowedByID,
	}, nil
}
