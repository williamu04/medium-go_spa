package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type UserRepository interface {
	SaveOneUser(ctx context.Context, user *model.UserModel) error
	FindOneUser(ctx context.Context, filter map[string]any) (*model.UserModel, error)
	FindAllUsers(ctx context.Context, filter map[string]any) ([]*model.UserModel, error)
	UpdateUser(ctx context.Context, user *model.UserModel, id uint) error
	DeleteUser(ctx context.Context, id uint) error
}
