package task

import (
	"bootcamp2/internal/core/port"
	"bootcamp2/internal/core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	router *gin.RouterGroup
	task   port.TaskService
}

func New(router *gin.RouterGroup, service *service.TaskService) *Handler {
	return &Handler{
		router: router,
		task:   service,
	}
}

// api/v1/task
func (h *Handler) RegisterRoutes() {
	group := h.router.Group("/task")
	{
		group.GET("/getAll", h.GetTaskList)
		group.GET("/get/:id", h.GetTask)
		group.POST("/create", h.InsertTask)
	}
}

func (h *Handler) GetTaskList(c *gin.Context) {
	result, err := h.task.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
func (h *Handler) GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get task",
	})
}
func (h *Handler) InsertTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "insert task",
	})

}
