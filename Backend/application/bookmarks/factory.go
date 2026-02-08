package bookmark_application

import "github.com/williamu04/medium-clone/domain/repository"

type BookmarkUseCase struct {
	Create      *CreateBookmarkUseCase
	Retrieve    *RetrieveBookmarkUseCase
	RetrieveAll *RetrieveAllBookmarkUseCase
	Delete      *DeleteBookmarkUseCase
}

func NewBookmarkUseCase(repository repository.BookmarkRepository) *BookmarkUseCase {
	return &BookmarkUseCase{
		Create:      NewCreateBookmarkUseCase(repository),
		Retrieve:    NewRetrieveBookmarkUseCase(repository),
		RetrieveAll: NewRetrieveAllBookmarkUseCase(repository),
		Delete:      NewDeleteBookmarkUseCase(repository),
	}
}
