package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type TopicRepository interface {
	SaveOneTopic(ctx context.Context, topic *model.TopicModel) error
	FindOneTopic(ctx context.Context, filter map[string]any) (*model.TopicModel, error)
	FindAllTopics(ctx context.Context, slugs []string) ([]*model.TopicModel, error)
	SetToString(ctx context.Context, topics []*model.TopicModel) ([]string, error)
	UpdateTopic(ctx context.Context, topic *model.TopicModel, id uint) error
	DeleteTopic(ctx context.Context, id uint) error
}
