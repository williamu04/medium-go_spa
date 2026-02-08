package user_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveUserUseCase struct {
	repository repository.UserRepository
}

func NewRetrieveUserUseCase(repository repository.UserRepository) *RetrieveUserUseCase {
	return &RetrieveUserUseCase{
		repository: repository,
	}
}

type RetrieveUserOutput struct {
	ID       uint
	Email    string
	Username string
	Bio      string
	Image    *string
}

func (uc *RetrieveUserUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveUserOutput, error) {
	user, err := uc.repository.FindOneUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return &RetrieveUserOutput{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
	}, nil
}
