package comment_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllCommentUseCase struct {
	repository repository.CommentRepository
}

func NewRetrieveAllCommentUseCase(repository repository.CommentRepository) *RetrieveAllCommentUseCase {
	return &RetrieveAllCommentUseCase{repository: repository}
}

type RetrieveAllCommentOutput struct {
	Comments []RetrieveCommentOutput
}

func (uc *RetrieveAllCommentUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveAllCommentOutput, error) {
	comments, err := uc.repository.FindAllComments(ctx, filter)

	if err != nil {
		return nil, err
	}

	if comments == nil {
		return nil, nil
	}

	return &RetrieveAllCommentOutput{
		Comments: mapCommentsToOutputs(comments),
	}, nil
}

func mapCommentsToOutputs(comments []*model.CommentModel) []RetrieveCommentOutput {
	outputs := make([]RetrieveCommentOutput, len(comments))

	for i, comment := range comments {
		outputs[i] = RetrieveCommentOutput{
			ID:        comment.ID,
			Body:      comment.Body,
			AuthorID:  comment.AuthorID,
			ArticleID: comment.ArticleID,
		}
	}
	return outputs
}
