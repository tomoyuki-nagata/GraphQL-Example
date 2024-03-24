package repository

import "errors"

var ErrCategoryNotFound = errors.New("category not found")

var ErrTaskNotFound = errors.New("task not found")

var ErrStatusInvalid = errors.New("invalid status")
