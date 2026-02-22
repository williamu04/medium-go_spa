package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type FollowRepository interface {
	SaveOneFollow(ctx context.Context, Follow *model.Follow) error
	FindOneFollow(ctx context.Context, filter map[string]any) (*model.Follow, error)
	FindAllFollows(ctx context.Context, filter map[string]any) ([]*model.Follow, error)
	DeleteFollow(ctx context.Context, id uint) error
}
