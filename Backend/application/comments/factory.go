package comment_application

import "github.com/williamu04/medium-clone/domain/repository"

type CommentUseCase struct {
	Create      *CreateCommentUseCase
	Retrieve    *RetrieveCommentUseCase
	RetrieveAll *RetrieveAllCommentUseCase
	Update      *UpdateCommentUseCase
	Delete      *DeleteCommentUseCase
}

func NewCommentUseCase(repository repository.CommentRepository) *CommentUseCase {
	return &CommentUseCase{
		Create:      NewCreateCommentUseCase(repository),
		Retrieve:    NewRetrieveCommentUseCase(repository),
		RetrieveAll: NewRetrieveAllCommentUseCase(repository),
		Update:      NewUpdateCommentUseCase(repository),
		Delete:      NewDeleteCommentUseCase(repository),
	}
}
