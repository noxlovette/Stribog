package handlers

import (
	"errors"
	"net/http"
	appError "stribog/internal/errors"
	"stribog/internal/services"
	"stribog/internal/types"

	"github.com/gin-gonic/gin"
)

type SparkHandler struct {
	Service *services.SparkService
}

func NewSparkHandler(service *services.SparkService) *SparkHandler {
	return &SparkHandler{Service: service}
}

func (h *SparkHandler) Get(c *gin.Context) {
	var SparkID = c.Param("id")

	ctx := c.Request.Context()
	spark, err := h.Service.GetSpark(ctx, SparkID)
	if err != nil {
		if errors.Is(err, appError.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spark not found"})
			return
		}
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, spark)
}

func (h *SparkHandler) Delete(c *gin.Context) {
	var SparkID = c.Param("id")

	ctx := c.Request.Context()
	err := h.Service.DeleteSpark(ctx, SparkID)
	if err != nil {
		if errors.Is(err, appError.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spark not found"})
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

func (h *SparkHandler) Update(c *gin.Context) {
	var SparkID = c.Param("id")

	var req types.SparkUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := c.Request.Context()
	err := h.Service.UpdateSpark(ctx, req, SparkID)
	if err != nil {
		if errors.Is(err, appError.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spark not found"})
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

func (h *SparkHandler) ListByForgeID(c *gin.Context) {
	forgeID := c.Param("forgeID")
	ctx := c.Request.Context()
	Sparks, err := h.Service.ListSparksByForgeID(ctx, forgeID)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Sparks)
}

func (h *SparkHandler) Create(c *gin.Context) {
	forgeID := c.Param("forgeID")
	ctx := c.Request.Context()
	SparkID, err := h.Service.CreateSpark(ctx, forgeID)
	if err != nil {
		if errors.Is(err, appError.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": SparkID})
}
