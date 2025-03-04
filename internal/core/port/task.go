package port

import (
	"bootcamp2/internal/core/domain/model"
	"context"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]*model.Task, error)
	GetById(ctx context.Context, params *model.TaskParams) (*model.Task, error)
	Create(ctx context.Context, task *model.TaskInsert) (*model.Task, error)
}
type TaskRepository interface {
}
