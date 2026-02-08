package dto

type TopicCreateDTO struct {
	Topic string `json:"topic" binding:"required,min=2"`
}

type TopicResponseDTO struct {
	ID    uint   `json:"topic_id"`
	Topic string `josn:"topic"`
	Slug  string `json:"topic_slug"`
}

type TopicUpdateDTO struct {
	Topic string `json:"topic" binding:"omitempty,min=2"`
}
