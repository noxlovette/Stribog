package handlers

import (
	"net/http"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
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

	c.SetCookie("access_token", tokens.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, 3600*24*7, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	accessToken, err := h.Service.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh failed"})
		return
	}

	c.SetCookie("access_token", accessToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "refreshed"})
}

func (h *UserHandler) Me(c *gin.Context) {
	userID, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
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
