package bookmarkapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveBookmarkUseCase struct {
	repository repository.BookmarkRepository
}

func NewRetrieveBookmarkUseCase(repository repository.BookmarkRepository) *RetrieveBookmarkUseCase {
	return &RetrieveBookmarkUseCase{repository: repository}
}

type RetrieveBookmarkOutput struct {
	ID        uint
	ArticleID uint
	UserID    uint
}

func (uc *RetrieveBookmarkUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveBookmarkOutput, error) {
	bookmark, err := uc.repository.FindOneBookmark(ctx, filter)

	if err != nil {
		return nil, err
	}

	if bookmark == nil {
		return nil, nil
	}

	return &RetrieveBookmarkOutput{
		ID:        bookmark.ID,
		ArticleID: bookmark.ArticleID,
		UserID:    bookmark.UserID,
	}, nil
}
