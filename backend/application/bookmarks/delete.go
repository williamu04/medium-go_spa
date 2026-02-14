package bookmarkapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type DeleteBookmarkUseCase struct {
	repository repository.BookmarkRepository
}

func NewDeleteBookmarkUseCase(repository repository.BookmarkRepository) *DeleteBookmarkUseCase {
	return &DeleteBookmarkUseCase{repository: repository}
}

func (uc *DeleteBookmarkUseCase) Execute(ctx context.Context, ID uint) error {
	if ID == 0 {
		return nil
	}

	Bookmark, err := uc.repository.FindOneBookmark(ctx, map[string]any{"id": ID})

	if err != nil {
		return err
	}

	if Bookmark == nil {
		return nil
	}

	return uc.repository.DeleteBookmark(ctx, ID)
}
