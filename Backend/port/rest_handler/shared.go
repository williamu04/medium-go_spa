package resthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamu04/medium-clone/pkg"
)

type SharedRestAPIHandler struct{}

func NewSharedRestAPIHandler() *SharedRestAPIHandler {
	return &SharedRestAPIHandler{}
}

func (h *SharedRestAPIHandler) RegisterSharedRoutes(router *gin.RouterGroup) {
	router.GET("/health", h.Health)
}

func (h *SharedRestAPIHandler) Health(c *gin.Context) {
	pkg.Success(c, http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "backend",
	})
}
