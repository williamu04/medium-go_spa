package userapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type UpdateUserUseCase struct {
	repository repository.UserRepository
	sluger     pkg.Sluger
}

func NewUpdateUserUseCase(repository repository.UserRepository, sluger pkg.Sluger) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		repository: repository,
		sluger:     sluger,
	}
}

type UpdateInput struct {
	Name     string
	Email    string
	Username string
	Bio      string
	Image    *string
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, input *UpdateInput, ID uint) error {
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

	if input.Name != "" {
		user.Name = input.Name
		user.Slug = uc.sluger.Slug(input.Name)
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Bio != "" {
		user.Bio = input.Bio
	}
	if input.Image != nil {
		user.Image = input.Image
	}

	return uc.repository.UpdateUser(ctx, user, ID)
}
