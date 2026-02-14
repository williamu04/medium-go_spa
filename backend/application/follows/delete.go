package followapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteFollowUseCase struct {
	repository repository.FollowRepository
}

func NewDeleteFollowUseCase(repository repository.FollowRepository) *DeleteFollowUseCase {
	return &DeleteFollowUseCase{repository: repository}
}

func (uc *DeleteFollowUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return nil
	}

	Follow, err := uc.repository.FindOneFollow(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if Follow == nil {
		return nil
	}

	return uc.repository.DeleteFollow(ctx, ID)
}
