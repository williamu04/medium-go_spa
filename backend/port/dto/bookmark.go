package dto

type BookmarkCreateDTO struct {
	ArticleID uint
	UserID    uint
}

type BookmarkResponseDTO struct {
	ID        uint
	ArticleID uint
	UserID    uint
}
