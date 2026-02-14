package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	followapplication "github.com/williamu04/medium-clone/application/follows"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type FollowRestAPIHandler struct {
	logger                   *pkg.Logger
	followCreateUseCase      *followapplication.CreateFollowUseCase
	followRetrieveUseCase    *followapplication.RetrieveFollowUseCase
	followRetrieveAllUseCase *followapplication.RetrieveAllFollowUseCase
	followDeleteUseCase      *followapplication.DeleteFollowUseCase
	authMiddleware           *middleware.AuthMiddleware
}

func NewfollowRestAPIHandler(
	logger *pkg.Logger,
	useCase *followapplication.FollowUseCase,
	auth *middleware.AuthMiddleware,
) *FollowRestAPIHandler {
	return &FollowRestAPIHandler{
		logger:                   logger,
		followCreateUseCase:      useCase.Create,
		followRetrieveUseCase:    useCase.Retrieve,
		followRetrieveAllUseCase: useCase.RetrieveAll,
		followDeleteUseCase:      useCase.Delete,
		authMiddleware:           auth,
	}
}

func (h *FollowRestAPIHandler) RegisterFollowRoutes(router *gin.RouterGroup) {
	router.GET("/:id", h.FollowRetrieve)
	router.GET("/all", h.FollowRetrieveAll)

	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())
	protected.POST("/create/:id", h.FollowCreate)
	protected.DELETE("/:id", h.FollowDelete)
}

func (h *FollowRestAPIHandler) FollowCreate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Retrieve: invalid Article ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Comment ID")
		return
	}

	user_id, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Follow ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	FollowedByID, ok := user_id.(uint)
	if !ok {
		h.logger.Error("Invalid Follow ID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	output, err := h.followCreateUseCase.Execute(c.Request.Context(), &followapplication.CreateFollowInput{
		FollowingID:  uint(id),
		FollowedByID: FollowedByID,
	})

	if err != nil {
		h.logger.Errorf("Creatioin failed for Follow")
	}

	res := dto.FollowResponseDTO{
		ID:           output.ID,
		FollowingID:  output.FollowingID,
		FollowedByID: output.FollowedByID,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *FollowRestAPIHandler) FollowRetrieve(c *gin.Context) {
	user_id, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Follow ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	FollowingID, ok := user_id.(uint)
	if !ok {
		h.logger.Error("Invalid Follow ID type")
		pkg.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.logger.Infof("Retrieve Follow attempt for ID: %d", FollowingID)

	output, err := h.followRetrieveUseCase.Execute(c.Request.Context(), map[string]any{"FollowingID": uint(FollowingID)})
	if err != nil {
		h.logger.Errorf("Retrieve Follow failed for ID %d: %v", FollowingID, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Follow")
		return
	}

	h.logger.Infof("Follow retrieved successfully: ID %d", FollowingID)

	res := dto.FollowResponseDTO{
		ID:           output.ID,
		FollowingID:  output.FollowingID,
		FollowedByID: output.FollowedByID,
	}

	pkg.Success(c, http.StatusOK, res)
}

func (h *FollowRestAPIHandler) FollowRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all Follows attempt")

	FollowingID, ok := c.Get("user_id")
	if !ok {
		h.logger.Error("Follow ID not found")
		pkg.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	output, err := h.followRetrieveAllUseCase.Execute(c.Request.Context(), map[string]any{"followingID": FollowingID.(uint)})
	if err != nil {
		h.logger.Errorf("Retrieve all Follows failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve Follows")
		return
	}

	h.logger.Infof("Follows retrieved successfully: %d Follows found", len(output.Follows))

	pkg.Success(c, http.StatusOK, output.Follows)
}

func (h *FollowRestAPIHandler) FollowDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid Follow ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid Follow ID")
		return
	}

	h.logger.Infof("Delete Follow attempt for ID: %d", id)

	if err := h.followDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete Follow failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete Follow")
		return
	}

	h.logger.Infof("Follow deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "Follow deleted successfully"})
}
