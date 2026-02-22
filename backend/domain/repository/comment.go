package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type CommentRepository interface {
	SaveOneComment(ctx context.Context, Comment *model.Comment) error
	FindOneComment(ctx context.Context, filter map[string]any) (*model.Comment, error)
	FindAllComments(ctx context.Context, filter map[string]any) ([]*model.Comment, error)
	UpdateComment(ctx context.Context, Comment *model.Comment, id uint) error
	DeleteComment(ctx context.Context, id uint) error
}
