package dataloader

import (
	"context"
	"errors"
	"graphql-example/graph/model"
	"graphql-example/repository"

	"github.com/graph-gophers/dataloader/v7"
)

type categoryBatcher struct {
	Repo repository.Repositories
}

// github.com/graph-gophers/dataloader/v7 の type BatchFunc[K, V]を満たすため
// dataloader.NewBatchedLoader関数の引数にできる
func (u *categoryBatcher) BatchGetCategories(ctx context.Context, ids []string) []*dataloader.Result[*model.Category] {
	// // 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[*model.Category], len(ids))
	for i := range results {
		results[i] = &dataloader.Result[*model.Category]{
			Error: errors.New("not found"),
		}
	}

	indexs := make(map[string]int, len(ids))
	for i, id := range ids {
		indexs[id] = i
	}

	categories, err := u.Repo.GetCategoryByIds(ctx, ids)
	for _, category := range categories {
		var rsl *dataloader.Result[*model.Category]
		if err != nil {
			rsl = &dataloader.Result[*model.Category]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.Category]{
				Data: category,
			}
		}
		// 該当するIDと同じindexに格納する
		results[indexs[category.ID]] = rsl
	}
	return results
}
