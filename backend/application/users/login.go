package userapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type LoginUserUseCase struct {
	repository     repository.UserRepository
	passwordHasher pkg.Hasher
	tokenGenerator pkg.JWTGen
}

func NewLoginUserUseCase(repository repository.UserRepository, passwordHasher pkg.Hasher, tokenGenerator pkg.JWTGen) *LoginUserUseCase {
	return &LoginUserUseCase{
		repository:     repository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Token string
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	if input.Username == "" || input.Password == "" {
		return nil, domain.ErrMissingFields
	}

	user, err := uc.repository.FindOneUser(ctx, map[string]any{"username": input.Username})

	if err != nil {
		return nil, domain.ErrInternalError
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	if err := uc.passwordHasher.Compare(user.PasswordHash, input.Password); err != nil {
		return nil, domain.ErrInvalidPassword
	}

	token, err := uc.tokenGenerator.Generate(user.ID, user.Email)

	if err != nil {
		return nil, domain.ErrInternalError
	}

	return &LoginOutput{
		Token: token,
	}, nil
}
