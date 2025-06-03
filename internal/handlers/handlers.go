package handlers

import (
	"net/http"
	"strconv"
	"taskproject/internal/middleware"
	"taskproject/internal/models"

	"github.com/gin-gonic/gin"
)

// Home Login - Handler
func (s *Server) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	token, err := middleware.CreateSession(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Get all tasks - Handler
func (s *Server) GetTasks(c *gin.Context) {
	usernameGet, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unauthenticated"})
		return
	}

	username := usernameGet.(string)

	var Tasks []models.Task
	err := s.db.Where("owner = ?", username).Find(&Tasks).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get tasks"})
		return
	}

	c.JSON(http.StatusOK, Tasks)
}

// Create one task - Handler
func (s *Server) PostTask(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unauthenticated"})
		return
	}

	var NewTask models.Task
	err := c.BindJSON(&NewTask)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "title required"})
		return
	}

	NewTask.Owner = username.(string)
	err = s.db.Create(&NewTask).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new task"})
		return
	}

	c.JSON(http.StatusCreated, NewTask)
}

// Get one task by ID - Handler
func (s *Server) GetTaskID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
		return
	}

	var task models.Task
	err = s.db.First(&task, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task ID not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Edit one task - Handler
func (s *Server) EditTask(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
		return
	}

	var UpdateTask models.Task
	err = s.db.First(&UpdateTask, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task ID not found"})
		return
	}

	err = c.BindJSON(&UpdateTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field to bind json"})
		return
	}

	err = s.db.Save(&UpdateTask).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save data"})
		return
	}

	c.JSON(http.StatusOK, UpdateTask)
}

// Delete task - Handler
func (s *Server) DeleleTaskID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
		return
	}

	var task models.Task
	err = s.db.Find(&task, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task ID not found"})
		return
	}

	// if exist then...
	err = s.db.Delete(&task, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
