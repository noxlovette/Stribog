package handlers

import (
	"errors"
	"net/http"
	appError "stribog/internal/errors"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
)

type AccessHandler struct {
	Service *services.AccessService
}

func NewAccessHandler(service *services.AccessService) *AccessHandler {
	return &AccessHandler{Service: service}
}

func (h *AccessHandler) List(c *gin.Context) {
	var forgeID = c.Param("forgeID")

	ctx := c.Request.Context()
	accessList, err := h.Service.ListForgeAccess(ctx, forgeID)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accessList)
}

func (h *AccessHandler) Delete(c *gin.Context) {
	var forgeID = c.Param("forgeID")
	var req types.AccessDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := c.Request.Context()
	err := h.Service.DeleteForgeAccess(ctx, forgeID, req)
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

func (h *AccessHandler) Create(c *gin.Context) {
	var req types.AccessCreateRequest
	var forgeID = c.Param("forgeID")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	err := h.Service.CreateForgeAccess(ctx, forgeID, req)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
