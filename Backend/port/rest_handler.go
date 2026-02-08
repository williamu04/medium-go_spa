package port

import (
	"github.com/gin-gonic/gin"
	article_application "github.com/williamu04/medium-clone/application/articles"
	bookmark_application "github.com/williamu04/medium-clone/application/bookmarks"
	comment_application "github.com/williamu04/medium-clone/application/comments"
	follow_application "github.com/williamu04/medium-clone/application/follows"
	topic_application "github.com/williamu04/medium-clone/application/topic"
	user_application "github.com/williamu04/medium-clone/application/users"
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
	user *user_application.UserUseCase,
	article *article_application.ArticleUseCase,
	topic *topic_application.TopicUseCase,
	comment *comment_application.CommentUseCase,
	bookmark *bookmark_application.BookmarkUseCase,
	follow *follow_application.FollowUseCase,
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
