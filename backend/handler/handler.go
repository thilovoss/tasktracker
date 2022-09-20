package handler

import (
	"errors"
	"net/http"
	"tasktracker/gormpatch"
	"tasktracker/models"
	"tasktracker/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group("/api")

	g.POST("/task", h.CreateTask)
	g.GET("/task", h.GetAllTasks)
	g.GET("/task/:id", h.GetTaskByID)
	g.PUT("/task/:id", h.UpdateTaskByID)
	g.DELETE("/task/:id", h.DeleteTask)

	g.POST("/category", h.CreateCategory)
	g.GET("/category", h.GetAllCategories)
	g.GET("/category/:id", h.GetCategoryById)
	g.PUT("/category/:id", h.UpdateCategoryByID)
	g.DELETE("/category/:id", h.DeleteCategory)
}

type IdInUriRequest struct {
	ID string `uri:"id" binding:"required"`
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

func (h *Handler) GetAllTasks(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)
	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
		return
	}

	task, err := services.GetTaskByID(uriReq.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTaskByID(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)
	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
	}

	var patches []gormpatch.JsonPatch
	err := c.BindJSON(&patches)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := services.PatchTask(uriReq.ID, &patches)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)

	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
		return
	}
	err := services.DeleteTask(uriReq.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}

type CategoryInput struct {
	Title string `json:"title" binding:"required"`
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var categoryInput CategoryInput

	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := &models.Category{
		Title: categoryInput.Title,
	}
	id, err := services.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetCategoryById(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)
	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
		return
	}

	category, err := services.GetCategoryById(uriReq.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) UpdateCategoryByID(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)
	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
	}

	var patches []gormpatch.JsonPatch
	err := c.BindJSON(&patches)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.PatchCategory(uriReq.ID, &patches)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	var uriReq *IdInUriRequest
	uriErr := c.ShouldBindUri(&uriReq)

	if uriErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uriErr.Error()})
		return
	}
	err := services.DeleteCategory(uriReq.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}
