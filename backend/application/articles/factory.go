package articleapplication

import (
	"github.com/williamu04/medium-clone/domain/repository"
	"github.com/williamu04/medium-clone/pkg"
)

type ArticleUseCase struct {
	Create      *CreateArticleUseCase
	Retrieve    *RetrieveArticleUseCase
	RetrieveAll *RetrieveAllArticleUseCase
	Update      *UpdateArticleUseCase
	Delete      *DeleteArticleUseCase
	Feed        *FeedArticleUseCase
}

func NewArticleUseCase(repository repository.ArticleRepository, topic repository.TopicRepository, user repository.UserRepository, comment repository.CommentRepository, bookmark repository.BookmarkRepository, sluger *pkg.Sluger) *ArticleUseCase {
	return &ArticleUseCase{
		Create:      NewCreateArticleUseCase(repository, topic, sluger),
		Retrieve:    NewRetrieveArticleUseCase(repository, topic),
		RetrieveAll: NewRetrieveAllArticlesUseCase(repository, topic),
		Update:      NewUpdateArticleUseCase(repository, topic, sluger),
		Delete:      NewDeleteArticleUseCase(repository),
		Feed:        NewFeedArticleUseCase(repository, user, comment, bookmark),
	}
}
