package service

import (
	"bootcamp2/internal/core/domain/model"
	"context"
)

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetAll(ctx context.Context) ([]*model.Task, error) {

}
func (s *TaskService) GetById(ctx context.Context, params *model.TaskParams) (*model.Task, error) {

}
func (s *TaskService) Create(ctx context.Context, task *model.TaskInsert) (*model.Task, error) {

}
