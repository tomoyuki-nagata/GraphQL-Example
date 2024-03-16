package repository

import (
	"context"
	"graphql-example/graph/model"
	"slices"
	"strconv"
)

var categories []*model.Category = []*model.Category{
	{ID: "v1_Category_1", Name: "仕事"},
	{ID: "v1_Category_2", Name: "趣味"},
	{ID: "v1_Category_3", Name: "勉強"},
}

var count int = len(categories)

type categoryRepository struct {
}

func (c *categoryRepository) GetAll(ctx context.Context) ([]*model.Category, error) {
	return categories, nil
}

func (c *categoryRepository) AddCategory(ctx context.Context, name string) (*model.Category, error) {
	count++
	newCategory := model.Category{
		ID:   "v1_Category_" + strconv.Itoa(count),
		Name: name,
	}
	categories = append(categories, &newCategory)
	return &newCategory, nil
}

func (c *categoryRepository) UpdateCategory(ctx context.Context, id, name string) (*model.Category, error) {
	for i, category := range categories {
		if category.ID == id {
			categories[i].Name = name
			return categories[i], nil
		}
	}
	return nil, ErrCategoryNotFound
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, id string) error {
	for i, category := range categories {
		if category.ID == id {
			categories = slices.Delete(categories, i, i+1)
			return nil
		}
	}
	return ErrCategoryNotFound
}
