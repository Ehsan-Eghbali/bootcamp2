package handler

import (
	"bootcamp2/internal/core/service"
	"bootcamp2/internal/presentation/handler/task"
	"github.com/gin-gonic/gin"
)

func Register(service *service.TaskService) {
	router := gin.Default()

	group := router.Group("/api/v1")

	task.New(group, service).RegisterRoutes()

	err := router.Run(":7070")
	if err != nil {
		panic(err)
	}

}
