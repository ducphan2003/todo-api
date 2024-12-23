package taskbiz

import (
	"context"
	"todo-api/common"
	"todo-api/database"
	taskModel "todo-api/modules/task/taskmodel"
)

type UpdateStorage interface {
	FindOneTask(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*taskModel.Task, error)
	UpdateTask(ctx context.Context,
		id uint,
		data *taskModel.TaskUpdate,
	) error
}

type updateBiz struct {
	updateStorage UpdateStorage
}

func NewUpdateBiz(updateStorage UpdateStorage) *updateBiz {
	return &updateBiz{
		updateStorage: updateStorage,
	}
}

func (biz *updateBiz) UpdateTask(ctx context.Context,
	id uint,
	data *taskModel.TaskUpdate,
) error {
	oldUser, err := biz.updateStorage.FindOneTask(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(taskModel.TaskEntityName, err)
	}
	if oldUser.Status == database.Deleted {
		return common.ErrEntityDeleted(taskModel.TaskEntityName, nil)
	}
	if err := biz.updateStorage.UpdateTask(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(taskModel.TaskEntityName, err)
	}
	return nil
}
