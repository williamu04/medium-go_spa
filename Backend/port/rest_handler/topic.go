package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	topic_application "github.com/williamu04/medium-clone/application/topic"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type TopicRestAPIHandler struct {
	logger                  *pkg.Logger
	topicCreateUseCase      *topic_application.CreateTopicUseCase
	topicRetrieveAllUseCase *topic_application.RetrieveAllTopicUseCase
	topicUpdateUseCase      *topic_application.UpdateTopicUseCase
	topicDeleteUseCase      *topic_application.DeleteTopicUseCase
	authMiddleware          *middleware.AuthMiddleware
}

func NewtopicRestAPIHandler(
	logger *pkg.Logger,
	useCase *topic_application.TopicUseCase,
	auth *middleware.AuthMiddleware,
) *TopicRestAPIHandler {
	return &TopicRestAPIHandler{
		logger:                  logger,
		topicCreateUseCase:      useCase.Create,
		topicRetrieveAllUseCase: useCase.RetrieveAll,
		topicUpdateUseCase:      useCase.Update,
		topicDeleteUseCase:      useCase.Delete,
		authMiddleware:          auth,
	}
}

func (h *TopicRestAPIHandler) RegisterTopicRoutes(router *gin.RouterGroup) {
	router.GET("/all", h.TopicRetrieveAll)

	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())
	protected.POST("/create", h.TopicCreate)
	protected.PUT("/:id", h.TopicUpdate)
	protected.DELETE("/:id", h.TopicDelete)
}

func (h *TopicRestAPIHandler) TopicCreate(c *gin.Context) {
	var req dto.TopicCreateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Create Topic: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	h.logger.Infof("Create attempt for Topic: %s", req.Topic)

	output, err := h.topicCreateUseCase.Execute(c.Request.Context(), &topic_application.CreateTopicInput{
		Topic: req.Topic,
	})

	if err != nil {
		h.logger.Errorf("Creatioin failed for Topic %s", req.Topic)
	}

	res := dto.TopicResponseDTO{
		ID:    output.ID,
		Topic: output.Topic,
		Slug:  output.Slug,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *TopicRestAPIHandler) TopicRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all Topics attempt")

	output, err := h.topicRetrieveAllUseCase.Execute(c.Request.Context(), nil)
	if err != nil {
		h.logger.Errorf("Retrieve all Topics failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Topics")
		return
	}

	h.logger.Infof("Topics retrieved successfully: %d Topics found", len(output.Topics))

	pkg.Success(c, http.StatusOK, output.Topics)
}

func (h *TopicRestAPIHandler) TopicUpdate(c *gin.Context) {
	idStr := c.Param("id")

	var req dto.TopicUpdateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Update: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Update: invalid Topic ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Topic ID")
		return
	}

	h.logger.Infof("Update Topic attempt for ID: %d with Topic: %s", id, req.Topic)

	if err := h.topicUpdateUseCase.Execute(c.Request.Context(), &topic_application.UpdateTopicInput{
		Topic: req.Topic,
	}, uint(id)); err != nil {
		h.logger.Errorf("Update Topic failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to update Topic")
		return
	}

	h.logger.Infof("Topic updated successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Topic updated successfully"})
}

func (h *TopicRestAPIHandler) TopicDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid Topic ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Topic ID")
		return
	}

	h.logger.Infof("Delete Topic attempt for ID: %d", id)

	if err := h.topicDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete Topic failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete Topic")
		return
	}

	h.logger.Infof("Topic deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Topic deleted successfully"})
}
