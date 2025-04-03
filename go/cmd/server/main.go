package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"stribog/config"

	"github.com/gin-gonic/gin"

	"stribog/internal/db/crud"
)

type Forge struct {
	ID string `uri:"id" binding:"required"`
}

func main() {
	ctx := context.Background()

	appState, err := config.InitAppState(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer appState.DB.Close()

	fmt.Println("Database connected successfully!")

	forgeService := crud.NewForgeService(appState.DB)

	route := gin.Default()

	route.GET("/:id", func(c *gin.Context) {
		var forge Forge
		if err := c.ShouldBindUri(&forge); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := forgeService.GetByID(c.Request.Context(), forge.ID)
		if err != nil {
			if errors.Is(err, crud.ErrForgeNotFound) {
				c.JSON(404, gin.H{"error": "Forge not found"})
				return
			}
			c.JSON(500, gin.H{"error": "Failed to fetch forge"})
			return
		}

		c.JSON(200, result)
	})

	route.Run(":8088")
}
