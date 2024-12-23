package taskstorage

import (
	"context"
	"todo-api/common"
	taskModel "todo-api/modules/task/taskmodel"
)

func (s *sqlStore) CreateTask(ctx context.Context, data *taskModel.TaskCreate) error {
	db := s.db.Begin()

	if err := db.Table(taskModel.Task{}.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
