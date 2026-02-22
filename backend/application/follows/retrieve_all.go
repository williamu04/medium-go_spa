package followapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllFollowUseCase struct {
	repository repository.FollowRepository
}

func NewRetrieveAllFollowUseCase(repository repository.FollowRepository) *RetrieveAllFollowUseCase {
	return &RetrieveAllFollowUseCase{repository: repository}
}

type RetrieveAllFollowOutput struct {
	Follows []RetrieveFollowOutput
}

func (uc *RetrieveAllFollowUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveAllFollowOutput, error) {
	Follows, err := uc.repository.FindAllFollows(ctx, filter)

	if err != nil {
		return nil, err
	}

	if Follows == nil {
		return nil, nil
	}

	return &RetrieveAllFollowOutput{
		Follows: mapFollowsToOutputs(Follows),
	}, nil
}

func mapFollowsToOutputs(Follows []*model.Follow) []RetrieveFollowOutput {
	outputs := make([]RetrieveFollowOutput, len(Follows))

	for i, Follow := range Follows {
		outputs[i] = RetrieveFollowOutput{
			ID:           Follow.ID,
			FollowingID:  Follow.FollowingID,
			FollowedByID: Follow.FollowedByID,
		}
	}
	return outputs
}
