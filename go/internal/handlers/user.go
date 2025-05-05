package handlers

import (
	"net/http"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service *services.UserService
	logger  *zap.Logger
}

func NewUserHandler(service *services.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		Service: service,
		logger:  logger,
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req types.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request body"})
		return
	}

	id, err := h.Service.RegisterUser(c.Request.Context(), req)
	if err != nil {
		switch err {
		case services.ErrPasswordTooShort:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case services.ErrEmailTaken:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": id.String()})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tokens, err := h.Service.Login(c.Request.Context(), req)
	if err != nil {
		switch err {
		case services.ErrEmailNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case services.ErrAuthenticationFailed:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, tokens)
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *UserHandler) Refresh(c *gin.Context) {
	h.logger.Info("Handling refresh token request")

	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.RefreshToken == "" {
		h.logger.Warn("Invalid refresh request", zap.Error(err), zap.String("token", req.RefreshToken))
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}

	h.logger.Info("Valid refresh token received, processing")

	tokens, err := h.Service.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		h.logger.Warn("Failed to refresh access token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh failed"})
		return
	}

	h.logger.Info("Access token refreshed successfully")
	c.JSON(http.StatusOK, tokens)
}

func (h *UserHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := h.Service.GetUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := h.Service.DeleteUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) Update(c *gin.Context) {
	var req types.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	err := h.Service.UpdateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
