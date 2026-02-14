package commentapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type UpdateCommentUseCase struct {
	repository repository.CommentRepository
}

func NewUpdateCommentUseCase(repository repository.CommentRepository) *UpdateCommentUseCase {
	return &UpdateCommentUseCase{repository: repository}
}

type UpdateCommentInput struct {
	Body string
}

func (uc *UpdateCommentUseCase) Execute(ctx context.Context, input *UpdateCommentInput, ID uint) error {
	if ID == 0 {
		return nil
	}

	comment, err := uc.repository.FindOneComment(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if input.Body != "" {
		comment.Body = input.Body
	}

	return uc.repository.UpdateComment(ctx, comment, ID)
}
