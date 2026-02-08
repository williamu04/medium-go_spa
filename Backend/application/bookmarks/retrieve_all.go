package bookmark_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllBookmarkUseCase struct {
	repository repository.BookmarkRepository
}

func NewRetrieveAllBookmarkUseCase(repository repository.BookmarkRepository) *RetrieveAllBookmarkUseCase {
	return &RetrieveAllBookmarkUseCase{repository: repository}
}

type RetrieveAllBookmarkOutput struct {
	Bookmarks []RetrieveBookmarkOutput
}

func (uc *RetrieveAllBookmarkUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveAllBookmarkOutput, error) {
	bookmarks, err := uc.repository.FindAllBookmarks(ctx, filter)

	if err != nil {
		return nil, err
	}

	if bookmarks == nil {
		return nil, nil
	}

	return &RetrieveAllBookmarkOutput{
		Bookmarks: mapBookmarksToOutputs(bookmarks),
	}, nil
}

func mapBookmarksToOutputs(bookmarks []*model.BookmarkModel) []RetrieveBookmarkOutput {
	outputs := make([]RetrieveBookmarkOutput, len(bookmarks))

	for i, bookmark := range bookmarks {
		outputs[i] = RetrieveBookmarkOutput{
			ID:        bookmark.ID,
			ArticleID: bookmark.ArticleID,
			UserID:    bookmark.UserID,
		}
	}
	return outputs
}
