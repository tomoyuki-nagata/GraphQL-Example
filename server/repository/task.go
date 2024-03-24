package repository

import (
	"context"
	"graphql-example/graph/model"
	"slices"
	"strconv"
	"time"
)

var tasks []*model.Task = []*model.Task{
	{ID: "v1_Task_1", Title: "仕事", Description: "", Category: &model.Category{ID: "v1_Category_1"}, Status: model.StatusWorking},
	{ID: "v1_Task_2", Title: "趣味", Description: "", Category: &model.Category{ID: "v1_Category_2"}, Status: model.StatusDone},
	{ID: "v1_Task_3", Title: "勉強", Description: "", Category: &model.Category{ID: "v1_Category_3"}, Status: model.StatusNew},
}

var tasksCount int = len(tasks)

type taskRepository struct {
}

func (c *taskRepository) GetAllTask(ctx context.Context) ([]*model.Task, error) {
	return tasks, nil
}

func (c *taskRepository) GetTaskById(ctx context.Context, id string) (*model.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, ErrTaskNotFound
}

func (c *taskRepository) CreateTask(ctx context.Context, title, description, categoryId string, deadline *time.Time) (*model.Task, error) {
	tasksCount++
	newTask := model.Task{
		ID:          "v1_Task_" + strconv.Itoa(tasksCount),
		Title:       title,
		Description: description,
		Category:    &model.Category{ID: categoryId},
		Status:      model.StatusNew,
	}
	if deadline != nil {
		newTask.Deadline = deadline
	}
	tasks = append(tasks, &newTask)
	return &newTask, nil
}

func (c *taskRepository) UpdateTask(ctx context.Context, id, title, description, categoryId, statusStr string, deadline *time.Time) (*model.Task, error) {
	status := model.Status(statusStr)
	for _, task := range tasks {
		if task.ID == id {
			task.Title = title
			task.Description = description
			task.Category = &model.Category{ID: categoryId}
			task.Status = model.Status(status)
			if deadline != nil {
				task.Deadline = deadline
			}
			return task, nil
		}
	}
	return nil, ErrTaskNotFound
}

func (c *taskRepository) DeleteTask(ctx context.Context, id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			return nil
		}
	}
	return ErrTaskNotFound
}
