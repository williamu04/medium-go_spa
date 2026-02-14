package dbrepository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewCommentDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DatabaseRepository) SaveOneComment(ctx context.Context, comment *model.CommentModel) error {
	if err := r.db.WithContext(ctx).Save(comment).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save comment %s: %v", comment.Body, err)
	}
	return nil
}

func (r *DatabaseRepository) FindOneComment(ctx context.Context, filter map[string]any) (*model.CommentModel, error) {
	var comment model.CommentModel

	if err := r.db.WithContext(ctx).Where(filter).First(&comment).Error; err != nil {
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

	return &comment, nil
}

func (r *DatabaseRepository) FindAllComments(ctx context.Context, filter map[string]any) ([]*model.CommentModel, error) {
	var comments []*model.CommentModel
	return comments, r.db.WithContext(ctx).Where(filter).Find(&comments).Error
}

func (r *DatabaseRepository) UpdateComment(ctx context.Context, comment *model.CommentModel, id uint) error {
	if err := r.db.WithContext(ctx).Model(comment).Where("id = ?", id).Updates(comment).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to update comment ID %d: %v", id, err)
	}
	return nil
}

func (r *DatabaseRepository) DeleteComment(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.CommentModel{}).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete comment %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Comment deleted %d", id)
	}
	return nil
}
