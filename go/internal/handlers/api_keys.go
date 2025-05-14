package handlers

import (
	"errors"
	"net/http"
	appError "stribog/internal/errors"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
)

type APIKeyHandler struct {
	Service *services.APIKeyService
}

func NewAPIKeyHandler(service *services.APIKeyService) *APIKeyHandler {
	return &APIKeyHandler{Service: service}
}

func (h *APIKeyHandler) List(c *gin.Context) {
	var forgeID = c.Param("forgeID")

	ctx := c.Request.Context()
	apiKeys, err := h.Service.ListAPIKeys(ctx, forgeID)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, apiKeys)
}

func (h *APIKeyHandler) Delete(c *gin.Context) {
	var forgeID = c.Param("forgeID")
	var req types.APIKeyID
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := c.Request.Context()
	err := h.Service.DeleteAPIKey(ctx, forgeID, req)
	if err != nil {
		if errors.Is(err, appError.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Forge not found"})
			return
		}
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *APIKeyHandler) Create(c *gin.Context) {
	var req types.CreateAPIKey
	var forgeID = c.Param("forgeID")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	apiKey, err := h.Service.CreateAPIKey(ctx, forgeID, req)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": apiKey})
}
