package repository

import (
	"context"
	"graphql-example/graph/model"
	"time"
)

type Repositories interface {
	TaskRepository
	CategoryRepository
}

type TaskRepository interface {
	GetAllTask(ctx context.Context) ([]*model.Task, error)
	GetTaskById(ctx context.Context, id string) (*model.Task, error)
	CreateTask(ctx context.Context, title, description, categoryId string, deadline *time.Time) (*model.Task, error)
	UpdateTask(ctx context.Context, id, title, description, categoryId, status string, deadline *time.Time) (*model.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) ([]*model.Category, error)
	GetCategoryByIds(ctx context.Context, ids []string) ([]*model.Category, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	AddCategory(ctx context.Context, name string) (*model.Category, error)
	UpdateCategory(ctx context.Context, id, name string) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type repositories struct {
	*taskRepository
	*categoryRepository
}

func New() Repositories {
	return &repositories{
		taskRepository:     &taskRepository{},
		categoryRepository: &categoryRepository{},
	}
}
