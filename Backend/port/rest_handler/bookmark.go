package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bookmark_application "github.com/williamu04/medium-clone/application/bookmarks"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type BookmarkRestAPIHandler struct {
	logger                     *pkg.Logger
	BookmarkCreateUseCase      *bookmark_application.CreateBookmarkUseCase
	BookmarkRetrieveUseCase    *bookmark_application.RetrieveBookmarkUseCase
	BookmarkRetrieveAllUseCase *bookmark_application.RetrieveAllBookmarkUseCase
	BookmarkDeleteUseCase      *bookmark_application.DeleteBookmarkUseCase
	authMiddleware             *middleware.AuthMiddleware
}

func NewBookmarkRestAPIHandler(
	logger *pkg.Logger,
	BookmarkUseCase *bookmark_application.BookmarkUseCase,
	auth *middleware.AuthMiddleware,
) *BookmarkRestAPIHandler {
	return &BookmarkRestAPIHandler{
		logger:                     logger,
		BookmarkCreateUseCase:      BookmarkUseCase.Create,
		BookmarkRetrieveUseCase:    BookmarkUseCase.Retrieve,
		BookmarkRetrieveAllUseCase: BookmarkUseCase.RetrieveAll,
		BookmarkDeleteUseCase:      BookmarkUseCase.Delete,
		authMiddleware:             auth,
	}
}

func (h *BookmarkRestAPIHandler) RegisterBookmarkRoutes(router *gin.RouterGroup) {
	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())

	protected.GET("/all", h.BookmarkRetrieveAll)
	protected.GET("/:id", h.BookmarkRetrieve)
	protected.POST("/create/:article_id", h.BookmarkCreate)
	protected.DELETE("/:id", h.BookmarkDelete)
}

func (h *BookmarkRestAPIHandler) BookmarkCreate(c *gin.Context) {
	idStr := c.Param("article_id")

	article_id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Retrieve: invalid Article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	user_id, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Bookmark ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := user_id.(uint)
	if !ok {
		h.logger.Error("Invalid Bookmark ID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	output, err := h.BookmarkCreateUseCase.Execute(c.Request.Context(), &bookmark_application.CreateBookmarkInput{
		UserID:    userID,
		ArticleID: uint(article_id),
	})

	if err != nil {
		h.logger.Errorf("Creatioin failed for Bookmark")
	}

	res := dto.BookmarkResponseDTO{
		ID:        output.ID,
		UserID:    output.UserID,
		ArticleID: output.ArticleID,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *BookmarkRestAPIHandler) BookmarkRetrieve(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Bookmark ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, ok := user_id.(uint)
	if !ok {
		h.logger.Error("Invalid Bookmark ID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.logger.Infof("Retrieve Bookmark attempt for ID: %d", userID)

	output, err := h.BookmarkRetrieveUseCase.Execute(c.Request.Context(), map[string]any{"user_id": uint(userID)})
	if err != nil {
		h.logger.Errorf("Retrieve Bookmark failed for ID %d: %v", userID, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Bookmark")
		return
	}

	h.logger.Infof("Bookmark retrieved successfully: ID %d", userID)

	res := dto.BookmarkResponseDTO{
		ID:        output.ID,
		UserID:    output.UserID,
		ArticleID: output.ArticleID,
	}

	pkg.Success(c, http.StatusOK, res)
}

func (h *BookmarkRestAPIHandler) BookmarkRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all Bookmarks attempt")

	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Bookmark ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	output, err := h.BookmarkRetrieveAllUseCase.Execute(c.Request.Context(), map[string]any{"userID": userID.(uint)})
	if err != nil {
		h.logger.Errorf("Retrieve all Bookmarks failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Bookmarks")
		return
	}

	h.logger.Infof("Bookmarks retrieved successfully: %d Bookmarks found", len(output.Bookmarks))

	pkg.Success(c, http.StatusOK, output.Bookmarks)
}

func (h *BookmarkRestAPIHandler) BookmarkDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid Bookmark ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Bookmark ID")
		return
	}

	h.logger.Infof("Delete Bookmark attempt for ID: %d", id)

	if err := h.BookmarkDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete Bookmark failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete Bookmark")
		return
	}

	h.logger.Infof("Bookmark deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Bookmark deleted successfully"})
}
