package bookmarkapplication

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type CreateBookmarkUseCase struct {
	repository repository.BookmarkRepository
}

func NewCreateBookmarkUseCase(repository repository.BookmarkRepository) *CreateBookmarkUseCase {
	return &CreateBookmarkUseCase{repository: repository}
}

type CreateBookmarkInput struct {
	ArticleID uint
	UserID    uint
}

type CreateBookmarkOutput struct {
	ID        uint
	ArticleID uint
	UserID    uint
}

func (uc *CreateBookmarkUseCase) Execute(ctx context.Context, input *CreateBookmarkInput) (*CreateBookmarkOutput, error) {
	bookmark := &model.Bookmark{
		ArticleID: input.ArticleID,
		UserID:    input.UserID,
	}

	if err := uc.repository.SaveOneBookmark(ctx, bookmark); err != nil {
		return nil, err
	}

	return &CreateBookmarkOutput{
		ID:        bookmark.ID,
		ArticleID: bookmark.ArticleID,
		UserID:    bookmark.UserID,
	}, nil
}
