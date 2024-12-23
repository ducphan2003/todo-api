package taskbiz

import (
	"context"
	"todo-api/common"
	taskModel "todo-api/modules/task/taskmodel"
)

type FindStorage interface {
	FindOneTask(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*taskModel.Task, error)
	ListTasksByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		filter *taskModel.TaskFilter,
		moreKeys ...string,
	) ([]taskModel.Task, error)
}

type findBiz struct {
	storage FindStorage
}

func NewFindBiz(storage FindStorage) *findBiz {
	return &findBiz{
		storage: storage,
	}
}

func (biz findBiz) ListTasksByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	filter *taskModel.TaskFilter,
	moreKeys ...string,
) ([]taskModel.Task, error) {
	moreKeys = append(moreKeys, "User")
	results, err := biz.storage.ListTasksByConditions(ctx, conditions, filter, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(taskModel.TaskEntityName, err)
	}
	return results, nil
}
