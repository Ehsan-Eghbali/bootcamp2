package bootstrap

import (
	"bootcamp2/internal/core/service"
	"bootcamp2/internal/presentation/handler"
	"context"
)

func Initiate(ctx context.Context) {
	taskSRV := service.NewTaskService()
	handler.Register(taskSRV)
}
