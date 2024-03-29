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

var categoryCount int = len(categories)

type categoryRepository struct {
}

func (c *categoryRepository) GetAllCategory(ctx context.Context) ([]*model.Category, error) {
	return categories, nil
}

func (c *categoryRepository) GetCategoryByIds(ctx context.Context, ids []string) ([]*model.Category, error) {
	result := []*model.Category{}
	for _, id := range ids {
		for _, category := range categories {
			if category.ID == id {
				result = append(result, category)
			}
		}
	}
	return result, nil
}

func (c *categoryRepository) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	for _, category := range categories {
		if category.ID == id {
			return category, nil
		}
	}
	return nil, ErrCategoryNotFound
}

func (c *categoryRepository) AddCategory(ctx context.Context, name string) (*model.Category, error) {
	categoryCount++
	newCategory := model.Category{
		ID:   "v1_Category_" + strconv.Itoa(categoryCount),
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
