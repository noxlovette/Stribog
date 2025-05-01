package main

import (
	"context"
	"log"
	"stribog/config"
	db "stribog/internal/db/sqlc"
	"stribog/internal/handlers"
	"stribog/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	state, err := config.InitAppState(ctx)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(state.DB.Pool)
	userService := services.NewUserService(queries)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()
	r.POST("/signup", userHandler.Signup)

	r.Run(":8080")
}
