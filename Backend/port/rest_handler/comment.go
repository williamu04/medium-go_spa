package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	comment_application "github.com/williamu04/medium-clone/application/comments"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type CommentRestAPIHandler struct {
	logger                    *pkg.Logger
	commentCreateUseCase      *comment_application.CreateCommentUseCase
	commentRetrieveUseCase    *comment_application.RetrieveCommentUseCase
	commentRetrieveAllUseCase *comment_application.RetrieveAllCommentUseCase
	commentUpdateUseCase      *comment_application.UpdateCommentUseCase
	commentDeleteUseCase      *comment_application.DeleteCommentUseCase
	authMiddleware            *middleware.AuthMiddleware
}

func NewcommentRestAPIHandler(
	logger *pkg.Logger,
	commentUseCase *comment_application.CommentUseCase,
	auth *middleware.AuthMiddleware,
) *CommentRestAPIHandler {
	return &CommentRestAPIHandler{
		logger:                    logger,
		commentCreateUseCase:      commentUseCase.Create,
		commentRetrieveUseCase:    commentUseCase.Retrieve,
		commentRetrieveAllUseCase: commentUseCase.RetrieveAll,
		commentUpdateUseCase:      commentUseCase.Update,
		commentDeleteUseCase:      commentUseCase.Delete,
		authMiddleware:            auth,
	}
}

func (h *CommentRestAPIHandler) RegisterCommentRoutes(router *gin.RouterGroup) {
	router.GET("/:id", h.CommentRetrieve)
	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())
	protected.POST("/create/:article_id", h.CommentCreate)
	protected.PUT("/:id", h.CommentUpdate)
	protected.DELETE("/:id", h.CommentDelete)
	protected.GET("/all", h.CommentRetrieveAll)
}

func (h *CommentRestAPIHandler) CommentCreate(c *gin.Context) {
	var req dto.CommentCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Create Comment: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	idStr := c.Param("article_id")

	article_id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Retrieve: invalid Article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	h.logger.Infof("Create attempt for Comment: %s", req.Body)

	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("User ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	authorID, ok := userID.(uint)
	if !ok {
		h.logger.Error("Invalid Comment ID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	output, err := h.commentCreateUseCase.Execute(c.Request.Context(), &comment_application.CreateCommentInput{
		Body:      req.Body,
		AuthorID:  authorID,
		ArticleID: uint(article_id),
	})

	if err != nil {
		h.logger.Errorf("Creatioin failed for Comment %s", req.Body)
	}

	res := dto.CommentResponseDTO{
		ID:        output.ID,
		Body:      output.Body,
		AuthorID:  output.AuthorID,
		ArticleID: output.ArticleID,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *CommentRestAPIHandler) CommentRetrieve(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Retrieve: invalid Comment ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	h.logger.Infof("Retrieve Comment attempt for ID: %d", id)

	output, err := h.commentRetrieveUseCase.Execute(c.Request.Context(), map[string]any{"id": uint(id)})
	if err != nil {
		h.logger.Errorf("Retrieve Comment failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Comment")
		return
	}

	h.logger.Infof("Comment retrieved successfully: ID %d", id)

	res := dto.CommentResponseDTO{
		ID:        output.ID,
		Body:      output.Body,
		AuthorID:  output.AuthorID,
		ArticleID: output.ArticleID,
	}

	pkg.Success(c, http.StatusOK, res)
}

func (h *CommentRestAPIHandler) CommentRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all Comments attempt")

	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Comment ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	output, err := h.commentRetrieveAllUseCase.Execute(c.Request.Context(), map[string]any{"author_id": userID.(uint)})
	if err != nil {
		h.logger.Errorf("Retrieve all Comments failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Comments")
		return
	}

	h.logger.Infof("Comments retrieved successfully: %d Comments found", len(output.Comments))

	pkg.Success(c, http.StatusOK, output.Comments)
}

func (h *CommentRestAPIHandler) CommentUpdate(c *gin.Context) {
	idStr := c.Param("id")

	var req dto.CommentUpdateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Update: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Update: invalid Comment ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	h.logger.Infof("Update Comment attempt for ID: %d with Comment: %s", id, req.Body)

	if err := h.commentUpdateUseCase.Execute(c.Request.Context(), &comment_application.UpdateCommentInput{
		Body: req.Body,
	}, uint(id)); err != nil {
		h.logger.Errorf("Update Comment failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to update Comment")
		return
	}

	h.logger.Infof("Comment updated successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (h *CommentRestAPIHandler) CommentDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid Comment ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	h.logger.Infof("Delete Comment attempt for ID: %d", id)

	if err := h.commentDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete Comment failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete Comment")
		return
	}

	h.logger.Infof("Comment deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
