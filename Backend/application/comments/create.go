package comment_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type CreateCommentUseCase struct {
	repository repository.CommentRepository
}

func NewCreateCommentUseCase(
	repository repository.CommentRepository,
) *CreateCommentUseCase {
	return &CreateCommentUseCase{
		repository: repository,
	}
}

type CreateCommentInput struct {
	Body      string
	ArticleID uint
	AuthorID  uint
}

type CreateCommentOutput struct {
	ID        uint
	Body      string
	ArticleID uint
	AuthorID  uint
}

func (uc *CreateCommentUseCase) Execute(ctx context.Context, input *CreateCommentInput) (*CreateCommentOutput, error) {
	comment := &model.CommentModel{
		Body:      input.Body,
		AuthorID:  input.AuthorID,
		ArticleID: input.ArticleID,
	}

	if err := uc.repository.SaveOneComment(ctx, comment); err != nil {
		return nil, err
	}

	return &CreateCommentOutput{
		ID:        comment.ID,
		Body:      comment.Body,
		AuthorID:  comment.AuthorID,
		ArticleID: comment.ArticleID,
	}, nil
}
