package dbrepository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/pkg"
	"gorm.io/gorm"
)

func NewUserDatabaseRepository(db *gorm.DB, logger *pkg.Logger) *DatabaseRepository {
	return &DatabaseRepository{db: db, logger: logger}
}

func (r *DatabaseRepository) SaveOneUser(ctx context.Context, user *model.User) error {
	err := r.db.Save(user).Error
	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to save user %s: %v", user.Username, err)
	}
	return err
}

func (r *DatabaseRepository) FindOneUser(ctx context.Context, filter map[string]any) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where(filter).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			if r.logger != nil {
				r.logger.Debugf("User not found with filter: %v", filter)
			}
			return nil, nil
		}
		if r.logger != nil {
			r.logger.Errorf("Database error finding user with filter %v: %v", filter, err)
		}
		return nil, err
	}
	return &user, nil
}

func (r *DatabaseRepository) FindAllUsers(ctx context.Context, filter map[string]any) ([]*model.User, error) {
	var users []*model.User
	return users, r.db.Where(filter).Find(&users).Error
}

func (r *DatabaseRepository) UpdateUser(ctx context.Context, user *model.User, id uint) error {
	err := r.db.Model(user).Where("id = ?", id).Updates(user).Error
	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to update user ID %d: %v", id, err)
	}
	return err
}

func (r *DatabaseRepository) DeleteUser(ctx context.Context, id uint) error {
	err := r.db.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil && r.logger != nil {
		r.logger.Errorf("Failed to delete user ID %d: %v", id, err)
	} else if err == nil && r.logger != nil {
		r.logger.Debugf("User ID %d deleted", id)
	}
	return err
}
