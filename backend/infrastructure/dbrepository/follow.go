package dbrepository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewFollowDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{
		db:     db,
		logger: logger,
	}
}

func (r *DatabaseRepository) SaveOneFollow(ctx context.Context, Follow *model.FollowModel) error {
	if err := r.db.WithContext(ctx).Save(Follow).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save Follow : %v", err)
	}
	return nil
}

func (r *DatabaseRepository) FindOneFollow(ctx context.Context, filter map[string]any) (*model.FollowModel, error) {
	var Follow model.FollowModel

	if err := r.db.WithContext(ctx).Where(filter).First(&Follow).Error; err != nil {
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

	return &Follow, nil
}

func (r *DatabaseRepository) FindAllFollows(ctx context.Context, filter map[string]any) ([]*model.FollowModel, error) {
	var Follows []*model.FollowModel
	return Follows, r.db.WithContext(ctx).Where(filter).Find(&Follows).Error
}

func (r *DatabaseRepository) DeleteFollow(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.FollowModel{}).Error; err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete Follow %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Follow deleted %d", id)
	}
	return nil
}
