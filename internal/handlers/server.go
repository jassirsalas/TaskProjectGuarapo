package handlers

import (
	"fmt"
	"net/http"

	"taskproject/internal/database"

	"gorm.io/gorm"
)

type Server struct {
	port int
	db   *gorm.DB
}

func NewServer() *http.Server {
	db, err := database.InitDB()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize database: %v", err))
	}

	NewServer := &Server{
		port: 8080,
		db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", NewServer.port),
		Handler: NewServer.RegisterRoutes(),
	}

	return server
}
