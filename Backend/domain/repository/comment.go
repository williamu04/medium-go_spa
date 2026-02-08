package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type CommentRepository interface {
	SaveOneComment(ctx context.Context, Comment *model.CommentModel) error
	FindOneComment(ctx context.Context, filter map[string]any) (*model.CommentModel, error)
	FindAllComments(ctx context.Context, filter map[string]any) ([]*model.CommentModel, error)
	UpdateComment(ctx context.Context, Comment *model.CommentModel, id uint) error
	DeleteComment(ctx context.Context, id uint) error
}
