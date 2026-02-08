package user_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain"
	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllUsersUseCase struct {
	repository repository.UserRepository
}

func NewRetrieveAllUsersUseCase(repository repository.UserRepository) *RetrieveAllUsersUseCase {
	return &RetrieveAllUsersUseCase{
		repository: repository,
	}
}

type RetrieveAllUsersOutput struct {
	Users []RetrieveUserOutput
}

func (uc *RetrieveAllUsersUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveAllUsersOutput, error) {
	users, err := uc.repository.FindAllUsers(ctx, filter)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, domain.ErrUserNotFound
	}

	return &RetrieveAllUsersOutput{
		Users: mapUsersToOutputs(users),
	}, nil
}

func mapUsersToOutputs(users []*model.UserModel) []RetrieveUserOutput {
	outputs := make([]RetrieveUserOutput, len(users))
	for i, user := range users {
		outputs[i] = RetrieveUserOutput{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Bio:      user.Bio,
			Image:    user.Image,
		}
	}
	return outputs
}
