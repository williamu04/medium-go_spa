package port

import (
	"github.com/gin-gonic/gin"
	articleapplication "github.com/williamu04/medium-clone/application/articles"
	bookmarkapplication "github.com/williamu04/medium-clone/application/bookmarks"
	commentapplication "github.com/williamu04/medium-clone/application/comments"
	followapplication "github.com/williamu04/medium-clone/application/follows"
	topicapplication "github.com/williamu04/medium-clone/application/topic"
	userapplication "github.com/williamu04/medium-clone/application/users"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	resthandler "github.com/williamu04/medium-clone/port/rest_handler"
)

type RestAPIHandler struct {
	sharedRestAPIHandler   *resthandler.SharedRestAPIHandler
	userRestAPIHandler     *resthandler.UserRestAPIHandler
	articleRestAPIHandler  *resthandler.ArticleRestAPIHandler
	topicRestAPIHandler    *resthandler.TopicRestAPIHandler
	commentRestAPIHandler  *resthandler.CommentRestAPIHandler
	bookmarkRestAPIHandler *resthandler.BookmarkRestAPIHandler
	followRestAPIHandler   *resthandler.FollowRestAPIHandler
}

func NewRestAPIHandler(
	logger *pkg.Logger,
	user *userapplication.UserUseCase,
	article *articleapplication.ArticleUseCase,
	topic *topicapplication.TopicUseCase,
	comment *commentapplication.CommentUseCase,
	bookmark *bookmarkapplication.BookmarkUseCase,
	follow *followapplication.FollowUseCase,
	auth *middleware.AuthMiddleware,
) *RestAPIHandler {
	return &RestAPIHandler{
		sharedRestAPIHandler:   resthandler.NewSharedRestAPIHandler(),
		userRestAPIHandler:     resthandler.NewUserRestAPIHandler(logger, user, auth),
		articleRestAPIHandler:  resthandler.NewArticleRestAPIHandler(logger, article, auth),
		topicRestAPIHandler:    resthandler.NewtopicRestAPIHandler(logger, topic, auth),
		commentRestAPIHandler:  resthandler.NewcommentRestAPIHandler(logger, comment, auth),
		bookmarkRestAPIHandler: resthandler.NewBookmarkRestAPIHandler(logger, bookmark, auth),
		followRestAPIHandler:   resthandler.NewfollowRestAPIHandler(logger, follow, auth),
	}
}

func (h *RestAPIHandler) RegisterRoutes(router *gin.RouterGroup) {
	h.sharedRestAPIHandler.RegisterSharedRoutes(router.Group(""))
	h.userRestAPIHandler.RegisterUserRoutes(router.Group("/user"))
	h.articleRestAPIHandler.RegisterArticleRoutes(router.Group("/article"))
	h.topicRestAPIHandler.RegisterTopicRoutes(router.Group("/topic"))
	h.commentRestAPIHandler.RegisterCommentRoutes(router.Group("/comment"))
	h.bookmarkRestAPIHandler.RegisterBookmarkRoutes(router.Group("/bookmark"))
	h.followRestAPIHandler.RegisterFollowRoutes(router.Group("/follow"))
}
