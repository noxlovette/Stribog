package main

import (
	"context"
	"log"
	"stribog/config"
	"stribog/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	appState, err := config.InitAppState(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer appState.DB.Close()

	router := gin.Default()

	routes.RegisterForgeRoutes(router, appState.ForgeService)

	log.Println("Server starting on :8088")
	router.Run(":8088")
}
