package dataloader

import (
	"graphql-example/graph/model"
	"graphql-example/repository"

	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	CategoryLoader dataloader.Interface[string, *model.Category]
}

func NewLoaders(repo repository.Repositories) *Loaders {
	categoryBatcher := &categoryBatcher{Repo: repo}

	return &Loaders{
		// dataloader.Loader[string, *model.User]型
		CategoryLoader: dataloader.NewBatchedLoader[string, *model.Category](categoryBatcher.BatchGetCategories),
	}
}
