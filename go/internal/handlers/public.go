package handlers

import (
	"errors"
	"net/http"
	appError "stribog/internal/errors"
	"stribog/internal/services"

	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	Service *services.PublicService
}

func NewPublicHandler(service *services.PublicService) *PublicHandler {
	return &PublicHandler{Service: service}
}

func (h *PublicHandler) List(c *gin.Context) {
	forgeID := c.Param("forgeID")
	ctx := c.Request.Context()
	Sparks, err := h.Service.ListSparksPublic(ctx, forgeID)
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
