package userapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type RegistrationUseCase struct {
	repository     repository.UserRepository
	passwordHasher pkg.Hasher
	tokenGenerator pkg.JWTGen
	sluger         pkg.Sluger
}

func NewRegistrationUseCase(repository repository.UserRepository, passwordHasher pkg.Hasher, tokenGenerator pkg.JWTGen, sluger pkg.Sluger) *RegistrationUseCase {
	return &RegistrationUseCase{
		repository:     repository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
		sluger:         sluger,
	}
}

type RegistrationInput struct {
	Name     string
	Email    string
	Username string
	Password string
	Bio      string
	Image    *string
}

type RegistrationOutput struct {
	ID       uint
	Name     string
	Slug     string
	Email    string
	Username string
	Password string
	Bio      string
	Image    *string
	Token    string
}

func (uc *RegistrationUseCase) Execute(ctx context.Context, input *RegistrationInput) (*RegistrationOutput, error) {
	if input.Email == "" || input.Username == "" || input.Password == "" {
		return nil, domain.ErrMissingFields
	}
	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Username: input.Username,
		Bio:      input.Bio,
		Image:    input.Image,
	}

	if !user.IsEmailValid() {
		return nil, domain.ErrInvalidEmail
	}

	if !user.IsPasswordValid(input.Password) {
		return nil, domain.ErrInvalidPassword
	}

	exist, _ := uc.repository.FindOneUser(ctx, map[string]any{"email": input.Email})
	if exist != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	exist, _ = uc.repository.FindOneUser(ctx, map[string]any{"username": input.Username})
	if exist != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	hashedPassword, err := uc.passwordHasher.Hash(input.Password)
	if err != nil {
		return nil, domain.ErrInternalError
	}

	user.PasswordHash = hashedPassword
	user.Slug = uc.sluger.Slug(input.Name)

	err = uc.repository.SaveOneUser(ctx, user)

	token, err := uc.tokenGenerator.Generate(user.ID, user.Email)
	if err != nil {
		return nil, domain.ErrInternalError
	}

	return &RegistrationOutput{
		ID:       user.ID,
		Name:     user.Name,
		Slug:     user.Slug,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
	}, nil
}
