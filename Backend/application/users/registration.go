package user_application

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
}

func NewRegistrationUseCase(repository repository.UserRepository, passwordHasher pkg.Hasher, tokenGenerator pkg.JWTGen) *RegistrationUseCase {
	return &RegistrationUseCase{
		repository:     repository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

type RegistrationInput struct {
	Email    string
	Username string
	Password string
	Bio      string
	Image    *string
}

type RegistrationOutput struct {
	ID       uint
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
	user := &model.UserModel{
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

	err = uc.repository.SaveOneUser(ctx, user)

	token, err := uc.tokenGenerator.Generate(user.ID, user.Email)

	if err != nil {
		return nil, domain.ErrInternalError
	}

	return &RegistrationOutput{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
	}, nil
}
