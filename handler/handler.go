package handler

import (
	"net/http"
	"tasktracker/models"
	"tasktracker/services"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/api")

	g.GET("/task", h.GetTasks)
	g.POST("/task", h.CreateTask)
}

func (h *Handler) GetTasks(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

type TaskInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Start       string `json:"start" binding:"required"`
	End         string `json:"end" binding:"required"`
	CategoryID  string `json:"categoryId"`
}

func (h *Handler) CreateTask(c *gin.Context) {
	var taskInput TaskInput

	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, parseErr := time.Parse(time.RFC3339, taskInput.Start)
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": parseErr.Error()})
		return
	}
	endTime, parseErr := time.Parse(time.RFC3339, taskInput.End)
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": parseErr.Error()})
		return
	}
	task := &models.Task{
		Name:        taskInput.Name,
		Description: taskInput.Description,
		Start:       startTime,
		End:         endTime,
		CategoryID:  taskInput.CategoryID,
	}
	id, err := services.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
