package dto

type CommentCreateDTO struct {
	Body string `json:"body" binding:"required,min=1"`
}

type CommentResponseDTO struct {
	ID        uint   `json:"comment_id"`
	Body      string `json:"body"`
	AuthorID  uint   `json:"author_id"`
	ArticleID uint   `json:"article_id"`
}

type CommentUpdateDTO struct {
	Body string `json:"body" binding:"omitempty,min=1"`
}
