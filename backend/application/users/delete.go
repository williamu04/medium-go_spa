package userapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteUserUseCase struct {
	repository repository.UserRepository
}

func NewDeleteUserUseCase(repository repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		repository: repository,
	}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return domain.ErrMissingFields
	}

	user, err := uc.repository.FindOneUser(ctx, map[string]any{"id": ID})
	if err != nil {
		return err
	}

	if user == nil {
		return domain.ErrUserNotFound
	}

	return uc.repository.DeleteUser(ctx, ID)
}
