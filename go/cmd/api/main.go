package main

import (
	"context"
	"log"
	"stribog/config"
	db "stribog/internal/db/sqlc"
	"stribog/internal/handlers"
	"stribog/internal/middleware"
	"stribog/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	state, err := config.InitAppState(ctx)
	if err != nil {
		log.Fatal(err)
	}

	querier := db.New(state.DB.Pool)
	userService := services.NewUserService(querier, state.TokenService, state.Logger)
	forgeService := services.NewForgeService(querier)
	accessService := services.NewAccessService(querier)
	sparkService := services.NewSparkService(querier)
	apiKeyService := services.NewAPIKeyService(querier)
	publicService := services.NewPublicService(querier)

	userHandler := handlers.NewUserHandler(userService, state.Logger)
	forgeHandler := handlers.NewForgeHandler(forgeService)
	accessHandler := handlers.NewAccessHandler(accessService)
	sparkHandler := handlers.NewSparkHandler(sparkService)
	apiKeyHandler := handlers.NewAPIKeyHandler(apiKeyService)
	publicHandler := handlers.NewPublicHandler(publicService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	authRoutes := r.Group("/auth")
	authRoutes.POST("/signup", userHandler.Signup)
	authRoutes.POST("/login", userHandler.Login)
	authRoutes.POST("/refresh", userHandler.Refresh)

	publicRoutes := r.Group("public")
	publicRoutes.Use(middleware.APIKeyAuthMiddleware(apiKeyService))

	publicRoutes.GET("/:forgeID", publicHandler.List)

	apiRoutes := r.Group("/api")
	apiRoutes.Use(middleware.AuthMiddlewareHeader(state.TokenService))

	userRoutes := apiRoutes.Group("/me")
	userRoutes.GET("/", userHandler.Get).DELETE("/", userHandler.Delete).PATCH("/", userHandler.Update)

	forgeRoutes := apiRoutes.Group("/forge")
	forgeRoutes.POST("/", forgeHandler.Create).GET("/", forgeHandler.List)
	forgeRoutes.GET("/:forgeID", forgeHandler.Get).DELETE("/:forgeID", forgeHandler.Delete).PATCH("/:forgeID", forgeHandler.Update)

	forgeSparks := forgeRoutes.Group("/:forgeID/sparks")
	forgeSparks.POST("/", sparkHandler.Create)
	forgeSparks.GET("/", sparkHandler.ListByForgeID)

	sparkRoutes := apiRoutes.Group("/sparks")
	sparkRoutes.GET("/:sparkID", sparkHandler.Get)
	sparkRoutes.PATCH("/:sparkID", sparkHandler.Update)
	sparkRoutes.DELETE("/:sparkID", sparkHandler.Delete)

	accessRoutes := forgeRoutes.Group("/:forgeID/access")
	accessRoutes.POST("/", accessHandler.Create).DELETE("/", accessHandler.Delete).GET("/", accessHandler.List)

	apiKeyRoutes := forgeRoutes.Group("/:forgeID/api-keys")
	apiKeyRoutes.POST("/", apiKeyHandler.Create).GET("/", apiKeyHandler.List).DELETE("/", apiKeyHandler.Delete)

	r.Run(":3000")
}
