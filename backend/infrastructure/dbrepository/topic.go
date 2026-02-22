package dbrepository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewTopicDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{db: db, logger: logger}
}

func (r *DatabaseRepository) SaveOneTopic(ctx context.Context, topic *model.Topic) error {
	err := r.db.Save(topic).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save Topic %s: %v", topic.Topic, err)
	}
	r.logger.Debugf("Row affected: %s", r.db.RowsAffected)
	return err
}

func (r *DatabaseRepository) FindOneTopic(ctx context.Context, filter map[string]any) (*model.Topic, error) {
	var Topic model.Topic

	err := r.db.WithContext(ctx).Where(filter).First(&Topic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if r.logger != nil {
				r.logger.Debugf("Topic not found with filter: %v", filter)
			}
			return nil, nil
		}
		if r.logger != nil {
			r.logger.Errorf("Database error finding Topic with filter %v: %v", filter, err)
		}
		return nil, err
	}
	return &Topic, nil
}

func (r *DatabaseRepository) FindAllTopics(ctx context.Context, slugs []string) ([]*model.Topic, error) {
	var Topics []*model.Topic

	if len(slugs) > 0 {
		return Topics, r.db.WithContext(ctx).Where("slug IN ?", slugs).Find(&Topics).Error
	}

	return Topics, r.db.WithContext(ctx).Find(&Topics).Error
}

func (r *DatabaseRepository) SetToString(ctx context.Context, topics []*model.Topic) ([]string, error) {
	slugs := make([]string, len(topics))

	for i, topic := range topics {
		slugs[i] = topic.Topic
	}

	return slugs, nil
}

func (r *DatabaseRepository) UpdateTopic(ctx context.Context, topic *model.Topic, id uint) error {
	err := r.db.Model(topic).Where("id = ?", id).Updates(topic).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to update Topic ID %d: %v", id, err)
	}
	return err
}

func (r *DatabaseRepository) DeleteTopic(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.Topic{}).Error

	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete Topic ID %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("Topic ID %d deleted", id)
	}
	return err
}
