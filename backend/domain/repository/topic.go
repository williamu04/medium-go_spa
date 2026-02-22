package repository

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
)

type TopicRepository interface {
	SaveOneTopic(ctx context.Context, topic *model.Topic) error
	FindOneTopic(ctx context.Context, filter map[string]any) (*model.Topic, error)
	FindAllTopics(ctx context.Context, slugs []string) ([]*model.Topic, error)
	SetToString(ctx context.Context, topics []*model.Topic) ([]string, error)
	UpdateTopic(ctx context.Context, topic *model.Topic, id uint) error
	DeleteTopic(ctx context.Context, id uint) error
}
