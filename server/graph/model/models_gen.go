// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Error interface {
	IsError()
	GetMessage() string
}

type Node interface {
	IsNode()
	GetID() string
}

type AddCategoryInput struct {
	Name string `json:"name"`
}

type AddCategoryPayload struct {
	Category *Category `json:"category,omitempty"`
	Errors   []Error   `json:"errors,omitempty"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Category) IsNode()            {}
func (this Category) GetID() string { return this.ID }

type CreateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  string     `json:"categoryId"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

type CreateTaskPayload struct {
	Task   *Task   `json:"task,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type DeleteCategoryPayload struct {
	ID     string  `json:"id"`
	Errors []Error `json:"errors,omitempty"`
}

type DeleteTaskPayload struct {
	ID     string  `json:"id"`
	Errors []Error `json:"errors,omitempty"`
}

type Mutation struct {
}

type NotFoundError struct {
	Message string `json:"message"`
}

func (NotFoundError) IsError()                {}
func (this NotFoundError) GetMessage() string { return this.Message }

type Query struct {
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    *Category  `json:"category"`
	Status      Status     `json:"status"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

func (Task) IsNode()            {}
func (this Task) GetID() string { return this.ID }

type UpdateCategoryInput struct {
	Name string `json:"name"`
}

type UpdateCategoryPayload struct {
	Category *Category `json:"category,omitempty"`
	Errors   []Error   `json:"errors,omitempty"`
}

type UpdateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  string     `json:"categoryId"`
	Status      Status     `json:"status"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

type UpdateTaskPayload struct {
	Task   *Task   `json:"task,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type Status string

const (
	StatusNew     Status = "NEW"
	StatusWorking Status = "WORKING"
	StatusDone    Status = "DONE"
)

var AllStatus = []Status{
	StatusNew,
	StatusWorking,
	StatusDone,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusNew, StatusWorking, StatusDone:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
