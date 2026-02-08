package comment_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveCommentUseCase struct {
	repository repository.CommentRepository
}

func NewRetrieveCommentUseCase(repository repository.CommentRepository) *RetrieveCommentUseCase {
	return &RetrieveCommentUseCase{repository: repository}
}

type RetrieveCommentOutput struct {
	ID        uint
	Body      string
	AuthorID  uint
	ArticleID uint
}

func (uc *RetrieveCommentUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveCommentOutput, error) {
	comment, err := uc.repository.FindOneComment(ctx, filter)
	if err != nil {
		return nil, err
	}

	if comment == nil {
		return nil, nil
	}

	return &RetrieveCommentOutput{
		ID:        comment.ID,
		Body:      comment.Body,
		AuthorID:  comment.AuthorID,
		ArticleID: comment.ArticleID,
	}, nil
}
