package handlers

import (
	"errors"
	"net/http"
	appError "stribog/internal/errors"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
)

type ForgeHandler struct {
	Service *services.ForgeService
}

func NewForgeHandler(service *services.ForgeService) *ForgeHandler {
	return &ForgeHandler{Service: service}
}

func (h *ForgeHandler) Get(c *gin.Context) {
	var forgeID = c.Param("id")

	ctx := c.Request.Context()
	user, err := h.Service.GetForge(ctx, forgeID)
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
	c.JSON(http.StatusOK, user)
}

func (h *ForgeHandler) Delete(c *gin.Context) {
	var forgeID = c.Param("id")

	ctx := c.Request.Context()
	err := h.Service.DeleteForge(ctx, forgeID)
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

func (h *ForgeHandler) Update(c *gin.Context) {
	var forgeID = c.Param("id")

	var req types.ForgeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	err := h.Service.UpdateForge(ctx, req, forgeID)
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

func (h *ForgeHandler) List(c *gin.Context) {
	ctx := c.Request.Context()
	forges, err := h.Service.ListForges(ctx)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, forges)
}

func (h *ForgeHandler) Create(c *gin.Context) {
	var req types.ForgeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	forgeID, err := h.Service.CreateForge(ctx, req)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": forgeID})
}
