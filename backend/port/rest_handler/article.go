package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	articleapplication "github.com/williamu04/medium-clone/application/articles"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type ArticleRestAPIHandler struct {
	logger                    *pkg.Logger
	articleCreateUseCase      *articleapplication.CreateArticleUseCase
	articleRetrieveUseCase    *articleapplication.RetrieveArticleUseCase
	articleRetrieveAllUseCase *articleapplication.RetrieveAllArticleUseCase
	articleUpdateUseCase      *articleapplication.UpdateArticleUseCase
	articleDeleteUseCase      *articleapplication.DeleteArticleUseCase
	authMiddleware            *middleware.AuthMiddleware
}

func NewArticleRestAPIHandler(
	logger *pkg.Logger,
	useCase *articleapplication.ArticleUseCase,
	auth *middleware.AuthMiddleware,
) *ArticleRestAPIHandler {
	return &ArticleRestAPIHandler{
		logger:                    logger,
		articleCreateUseCase:      useCase.Create,
		articleRetrieveUseCase:    useCase.Retrieve,
		articleRetrieveAllUseCase: useCase.RetrieveAll,
		articleUpdateUseCase:      useCase.Update,
		articleDeleteUseCase:      useCase.Delete,
		authMiddleware:            auth,
	}
}

func (h *ArticleRestAPIHandler) RegisterArticleRoutes(router *gin.RouterGroup) {
	router.GET("/:id", h.ArticleRetrieve)

	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())

	protected.GET("/all", h.ArticleRetrieveAll)
	protected.POST("/create", h.ArticleCreate)
	protected.PUT("/:id", h.ArticleUpdate)
	protected.DELETE("/:id", h.ArticleDelete)
}

func (h *ArticleRestAPIHandler) ArticleCreate(c *gin.Context) {
	var req dto.ArticleCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Create Article: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	h.logger.Infof("Create attempt for article: %s", req.Title)

	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("User ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	authorID, ok := userID.(uint)
	if !ok {
		h.logger.Error("Invalid authorID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	output, err := h.articleCreateUseCase.Execute(c.Request.Context(), &articleapplication.CreateArticleInput{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		Topic:       req.Topic,
		AuthorID:    authorID,
	})

	if err != nil || output == nil {
		h.logger.Errorf("Creatioin failed for article %s", req.Title)
		pkg.Error(c, http.StatusInternalServerError, "Article not created")
	}

	res := dto.ArticleResponseDTO{
		ID:          output.ID,
		Title:       output.Title,
		Slug:        output.Slug,
		Description: output.Description,
		Body:        output.Body,
		AuthorID:    output.AuthorID,
		Topic:       output.Topic,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *ArticleRestAPIHandler) ArticleRetrieve(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Retrieve: invalid article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	h.logger.Infof("Retrieve article attempt for ID: %d", id)

	output, err := h.articleRetrieveUseCase.Execute(c.Request.Context(), map[string]any{"id": uint(id)})
	if err != nil {
		h.logger.Errorf("Retrieve article failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve article")
		return
	}

	h.logger.Infof("article retrieved successfully: ID %d", id)

	res := dto.ArticleResponseDTO{
		ID:          output.ID,
		Title:       output.Title,
		Slug:        output.Slug,
		Description: output.Description,
		Body:        output.Body,
		AuthorID:    output.AuthorID,
		Topic:       output.Topic,
	}

	pkg.Success(c, http.StatusOK, res)
}

func (h *ArticleRestAPIHandler) ArticleRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all articles attempt")

	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("User ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	output, err := h.articleRetrieveAllUseCase.Execute(c.Request.Context(), map[string]any{"authorID": userID.(uint)})
	if err != nil {
		h.logger.Errorf("Retrieve all articles failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve articles")
		return
	}

	h.logger.Infof("articles retrieved successfully: %d articles found", len(output.Articles))

	pkg.Success(c, http.StatusOK, output.Articles)
}

func (h *ArticleRestAPIHandler) ArticleUpdate(c *gin.Context) {
	idStr := c.Param("id")

	var req dto.ArticleUpdateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Update: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Update: invalid article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	h.logger.Infof("Update article attempt for ID: %d with article: %s", id, req.Title)

	if err := h.articleUpdateUseCase.Execute(c.Request.Context(), &articleapplication.UpdateArticleInput{
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		Topic:       req.Topic,
	}, uint(id)); err != nil {
		h.logger.Errorf("Update article failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to update article")
		return
	}

	h.logger.Infof("article updated successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "article updated successfully"})
}

func (h *ArticleRestAPIHandler) ArticleDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid article ID")
		return
	}

	h.logger.Infof("Delete article attempt for ID: %d", id)

	if err := h.articleDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete article failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete article")
		return
	}

	h.logger.Infof("article deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "article deleted successfully"})
}
