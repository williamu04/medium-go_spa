package dto

type ArticleCreateDTO struct {
	Title       string   `json:"title" binding:"required,min=3"`
	Description string   `json:"description" binding:"required,min=8"`
	Body        string   `json:"body" binding:"required"`
	Topic       []string `json:"topic" binding:"required,min=1"`
}

type ArticleResponseDTO struct {
	ID          uint     `json:"article_id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	AuthorID    uint     `json:"author_id"`
	Topic       []string `json:"topic"`
}

type ArticleUpdateDTO struct {
	Title       string   `json:"title" binding:"omitempty,min=3"`
	Description string   `json:"description" binding:"omitempty,min=8"`
	Body        string   `json:"body" binding:"omitempty"`
	Topic       []string `json:"topic" binding:"omitempty,min=1"`
}

type ArticleFeed struct {
	ID            uint   `json:"article_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Thumbnail     string `json:"thumbnail_url"`
	AuthorImage   string `json:"author_image"`
	AuthorName    string `json:"author_name"`
	CommentCount  int    `json:"comment_count"`
	BookmarkCount int    `json:"bookmark_count"`
}
