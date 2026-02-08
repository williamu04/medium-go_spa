package resthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_application "github.com/williamu04/medium-clone/application/users"
	"github.com/williamu04/medium-clone/middleware"
	"github.com/williamu04/medium-clone/pkg"
	"github.com/williamu04/medium-clone/port/dto"
)

type UserRestAPIHandler struct {
	logger                  *pkg.Logger
	userRegistrationUseCase *user_application.RegistrationUseCase
	userLoginUseCase        *user_application.LoginUserUseCase
	userRetrieveUseCase     *user_application.RetrieveUserUseCase
	userRetrieveAllUseCase  *user_application.RetrieveAllUsersUseCase
	userUpdateUseCase       *user_application.UpdateUserUseCase
	userDeleteUseCase       *user_application.DeleteUserUseCase
	authMiddleware          *middleware.AuthMiddleware
}

func NewUserRestAPIHandler(
	logger *pkg.Logger,
	useCase *user_application.UserUseCase,
	auth *middleware.AuthMiddleware,
) *UserRestAPIHandler {

	return &UserRestAPIHandler{
		logger:                  logger,
		userRegistrationUseCase: useCase.Registration,
		userLoginUseCase:        useCase.Login,
		userRetrieveUseCase:     useCase.Retrieve,
		userRetrieveAllUseCase:  useCase.RetrieveAll,
		userUpdateUseCase:       useCase.Update,
		userDeleteUseCase:       useCase.Delete,
		authMiddleware:          auth,
	}
}

func (h *UserRestAPIHandler) RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.UserRegistration)
	router.POST("/login", h.UserLogin)
	router.GET("/:id", h.UserRetrieve)
	router.GET("/all", h.UserRetrieveAll)

	protected := router.Group("")
	protected.Use(h.authMiddleware.Auth())
	protected.PUT("/:id", h.UserUpdate)
	protected.DELETE("/:id", h.UserDelete)
}

func (h *UserRestAPIHandler) UserRegistration(c *gin.Context) {
	var req dto.UserRegistrationDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Registration: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	h.logger.Infof("Registration attempt for user: %s", req.Username)

	output, err := h.userRegistrationUseCase.Execute(c.Request.Context(), &user_application.RegistrationInput{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Bio:      req.Bio,
		Image:    &req.Image,
	})

	if err != nil {
		h.logger.Errorf("Registration failed for user %s: %v", req.Username, err)
		pkg.Error(c, http.StatusInternalServerError, "Registration failed")
		return
	}

	h.logger.Infof("User registered successfully: %s (ID: %d)", req.Username, output.ID)

	res := dto.UserResponseDTO{
		ID:       output.ID,
		Email:    output.Email,
		Username: output.Username,
		Bio:      output.Bio,
		Image:    *output.Image,
		Token:    output.Token,
	}

	pkg.Success(c, http.StatusCreated, res)
}

func (h *UserRestAPIHandler) UserLogin(c *gin.Context) {
	var req dto.UserLoginDTO

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		h.logger.Warnf("Login: Invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
	}

	h.logger.Infof("Login attempt for user: %s", req.Username)

	output, err := h.userLoginUseCase.Execute(c.Request.Context(), &user_application.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		h.logger.Errorf("Login failed for user %s: %v", req.Username, err)
		pkg.Error(c, http.StatusBadRequest, "Login failed")
	}

	h.logger.Infof("User login successfully: %s (ID: %d)", req.Username)

	res := dto.UserResponseDTO{
		Token: output.Token,
	}

	pkg.Success(c, http.StatusAccepted, res)
}

func (h *UserRestAPIHandler) UserRetrieve(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Retrieve: invalid user ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	h.logger.Infof("Retrieve user attempt for ID: %d", id)

	output, err := h.userRetrieveUseCase.Execute(c.Request.Context(), map[string]any{"id": uint(id)})
	if err != nil {
		h.logger.Errorf("Retrieve user failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	h.logger.Infof("User retrieved successfully: ID %d", id)

	res := dto.UserResponseDTO{
		ID:       output.ID,
		Email:    output.Email,
		Username: output.Username,
		Bio:      output.Bio,
		Image:    *output.Image,
	}

	pkg.Success(c, http.StatusOK, res)
}

func (h *UserRestAPIHandler) UserRetrieveAll(c *gin.Context) {
	h.logger.Infof("Retrieve all users attempt")

	output, err := h.userRetrieveAllUseCase.Execute(c.Request.Context(), nil)
	if err != nil {
		h.logger.Errorf("Retrieve all users failed: %v", err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	h.logger.Infof("Users retrieved successfully: %d users found", len(output.Users))

	pkg.Success(c, http.StatusOK, output.Users)
}

func (h *UserRestAPIHandler) UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	var req dto.UserUpdateDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warnf("Update: invalid request body - %v", err)
		pkg.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.Warnf("Update: invalid user ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	h.logger.Infof("Update user attempt for ID: %d with username: %s", id, req.Username)

	if err := h.userUpdateUseCase.Execute(c.Request.Context(), &user_application.UpdateInput{
		Email:    req.Email,
		Username: req.Username,
		Bio:      req.Bio,
		Image:    &req.Image,
	}, uint(id)); err != nil {
		h.logger.Errorf("Update user failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	h.logger.Infof("User updated successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserRestAPIHandler) UserDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Warnf("Delete: invalid user ID format - %s", idStr)
		pkg.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	h.logger.Infof("Delete user attempt for ID: %d", id)

	if err := h.userDeleteUseCase.Execute(c.Request.Context(), uint(id)); err != nil {
		h.logger.Errorf("Delete user failed for ID %d: %v", id, err)
		pkg.Error(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	h.logger.Infof("User deleted successfully: ID %d", id)

	pkg.Success(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
