package taskbiz

import (
	"context"
	"todo-api/common"
	"todo-api/database"
	taskModel "todo-api/modules/task/taskmodel"
)

type DeleteStorage interface {
	FindOneTask(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*taskModel.Task, error)
	DeleteTask(ctx context.Context,
		id uint,
	) error
}

type deleteBiz struct {
	storage DeleteStorage
}

func NewDeleteBiz(storage DeleteStorage) *deleteBiz {
	return &deleteBiz{
		storage: storage,
	}
}

func (biz deleteBiz) DeleteTask(ctx context.Context,
	id uint) error {
	oldData, err := biz.storage.FindOneTask(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(taskModel.TaskEntityName, err)
	}
	if oldData.Status == database.Deleted {
		return common.ErrEntityDeleted(taskModel.TaskEntityName, err)
	}
	if err := biz.storage.DeleteTask(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(taskModel.Task{}.TableName(), err)
	}
	return nil
}
