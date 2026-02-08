package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type BookmarkRepository interface {
	SaveOneBookmark(ctx context.Context, Bookmark *model.BookmarkModel) error
	FindOneBookmark(ctx context.Context, filter map[string]any) (*model.BookmarkModel, error)
	FindAllBookmarks(ctx context.Context, filter map[string]any) ([]*model.BookmarkModel, error)
	DeleteBookmark(ctx context.Context, id uint) error
}
