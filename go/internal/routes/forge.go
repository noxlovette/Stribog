package routes

import (
	"errors"
	"net/http"
	"stribog/internal/services"

	"github.com/gin-gonic/gin"
)

type Forge struct {
	ID string `uri:"id" binding:"required"`
}

func RegisterForgeRoutes(r *gin.Engine, forgeService *services.ForgeService) {
	r.GET("/:id", func(c *gin.Context) {
		var forge Forge
		if err := c.ShouldBindUri(&forge); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := forgeService.GetByID(c.Request.Context(), forge.ID)
		if err != nil {
			if errors.Is(err, services.ErrForgeNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Forge not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch forge"})
			return
		}

		c.JSON(http.StatusOK, result)
	})
}
