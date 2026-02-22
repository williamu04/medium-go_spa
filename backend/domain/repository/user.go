package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type UserRepository interface {
	SaveOneUser(ctx context.Context, user *model.User) error
	FindOneUser(ctx context.Context, filter map[string]any) (*model.User, error)
	FindAllUsers(ctx context.Context, filter map[string]any) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, id uint) error
	DeleteUser(ctx context.Context, id uint) error
}
