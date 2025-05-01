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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
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

	user, err := h.Service.Login(c.Request.Context(), req)
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

	c.JSON(http.StatusOK, gin.H{"user": user})
}
