package handlers

import (
	"net/http"
	"taskproject/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // not used, set for frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Home
	r.POST("/login", s.Login)

	// Route-specific middleware
	protectedRoutes := r.Group("/tasks")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		// Auth
		protectedRoutes.GET("", s.GetTasks)
		protectedRoutes.POST("", s.PostTask)
		protectedRoutes.GET("/:id", s.GetTaskID)
		protectedRoutes.PUT("/:id", s.EditTask)
		protectedRoutes.DELETE(":id", s.DeleleTaskID)
	}

	return r
}
