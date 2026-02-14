package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type FollowRepository interface {
	SaveOneFollow(ctx context.Context, Follow *model.FollowModel) error
	FindOneFollow(ctx context.Context, filter map[string]any) (*model.FollowModel, error)
	FindAllFollows(ctx context.Context, filter map[string]any) ([]*model.FollowModel, error)
	DeleteFollow(ctx context.Context, id uint) error
}
