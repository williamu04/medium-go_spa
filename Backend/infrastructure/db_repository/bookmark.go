package db_repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewBookmarkDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DatabaseRepository) SaveOneBookmark(ctx context.Context, Bookmark *model.BookmarkModel) error {
	if err := r.db.WithContext(ctx).Save(Bookmark).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save Bookmark : %v", err)
	}
	return nil
}

func (r *DatabaseRepository) FindOneBookmark(ctx context.Context, filter map[string]any) (*model.BookmarkModel, error) {
	var Bookmark model.BookmarkModel

	if err := r.db.WithContext(ctx).Where(filter).First(&Bookmark).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if r.logger != nil {
				r.logger.Debugf("Article not found with filter: %v", filter)
			}
			return nil, nil
		}
		if r.logger != nil {
			r.logger.Errorf("Database error finding article with filter %v: %v", filter, err)
		}
		return nil, err
	}

	return &Bookmark, nil
}

func (r *DatabaseRepository) FindAllBookmarks(ctx context.Context, filter map[string]any) ([]*model.BookmarkModel, error) {
	var Bookmarks []*model.BookmarkModel
	return Bookmarks, r.db.WithContext(ctx).Where(filter).Find(&Bookmarks).Error
}

func (r *DatabaseRepository) DeleteBookmark(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.BookmarkModel{}).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete Bookmark %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Bookmark deleted %d", id)
	}
	return nil
}
