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

	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	// Public auth routes
	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.POST("/refresh", userHandler.Refresh)

	// Protected routes
	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware(state.TokenService))
	authRoutes.GET("/me", userHandler.Me)

	r.Run(":8080")
}
