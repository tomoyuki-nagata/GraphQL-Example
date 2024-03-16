package repository

import (
	"context"
	"graphql-example/graph/model"
)

type Repositories interface {
	CategoryRepository
}

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]*model.Category, error)
	AddCategory(ctx context.Context, name string) (*model.Category, error)
	UpdateCategory(ctx context.Context, id, name string) (*model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type repositories struct {
	*categoryRepository
}

func New() Repositories {
	return &repositories{
		categoryRepository: &categoryRepository{},
	}
}
