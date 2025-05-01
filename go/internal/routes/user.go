package routes

import (
	db "stribog/internal/db/sqlc"
	"stribog/internal/handlers"
	"stribog/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(queries *db.Queries) *gin.Engine {
	r := gin.Default()

	userService := services.NewUserService(queries)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/signup", userHandler.Signup)

	return r
}
