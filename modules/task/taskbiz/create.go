package taskbiz

import (
	"context"
	"todo-api/common"
	taskModel "todo-api/modules/task/taskmodel"
)

type CreateStorage interface {
	CreateTask(ctx context.Context, data *taskModel.TaskCreate) error
}

type createBiz struct {
	createStorage CreateStorage
}

func NewCreateBiz(createStorage CreateStorage) *createBiz {
	return &createBiz{
		createStorage: createStorage,
	}
}

func (biz createBiz) CreateTask(ctx context.Context, data *taskModel.TaskCreate) error {
	if err := biz.createStorage.CreateTask(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(taskModel.Task{}.TableName(), err)
	}
	return nil
}
