package commentapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteCommentUseCase struct {
	repository repository.CommentRepository
}

func NewDeleteCommentUseCase(repository repository.CommentRepository) *DeleteCommentUseCase {
	return &DeleteCommentUseCase{repository: repository}
}

func (uc *DeleteCommentUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return nil
	}

	Comment, err := uc.repository.FindOneComment(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if Comment == nil {
		return nil
	}

	return uc.repository.DeleteComment(ctx, ID)
}
