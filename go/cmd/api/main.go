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
	userService := services.NewUserService(querier, state.TokenService)
	forgeService := services.NewForgeService(querier)
	accessService := services.NewAccessService(querier)
	sparkService := services.NewSparkService(querier)

	userHandler := handlers.NewUserHandler(userService)
	forgeHandler := handlers.NewForgeHandler(forgeService)
	accessHandler := handlers.NewAccessHandler(accessService)
	sparkHandler := handlers.NewSparkHandler(sparkService)

	r := gin.Default()

	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.POST("/refresh", userHandler.Refresh)

	apiRoutes := r.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware(state.TokenService))
	apiRoutes.GET("/me", userHandler.Me)

	userRoutes := apiRoutes.Group("/user")
	userRoutes.GET("/", userHandler.Get).DELETE("/", userHandler.Delete).PATCH("/", userHandler.Update)

	sparkRoutes := apiRoutes.Group("/spark")
	sparkRoutes.POST("/new/:forgeID", sparkHandler.Create).GET("/list/:forgeID", sparkHandler.ListByForgeID)
	sparkRoutes.GET("/:id", sparkHandler.Get).DELETE("/:id", sparkHandler.Delete).PATCH("/:id", sparkHandler.Update)

	forgeRoutes := apiRoutes.Group("/forge")
	forgeRoutes.POST("/", forgeHandler.Create).GET("/all", forgeHandler.List)
	forgeRoutes.GET("/:id", forgeHandler.Get).DELETE("/:id", forgeHandler.Delete).PATCH("/:id", forgeHandler.Update)

	accessRoutes := forgeRoutes.Group("/:id/access")
	accessRoutes.POST("/", accessHandler.Create).DELETE("/", accessHandler.Delete).GET("/", accessHandler.List)

	r.Run(":8080")
}
