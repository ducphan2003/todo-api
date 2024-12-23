package taskstorage

import (
	"context"
	"todo-api/common"
	"todo-api/database"
	taskModel "todo-api/modules/task/taskmodel"
)

func (s *sqlStore) DeleteTask(ctx context.Context,
	id uint,
) error {
	db := s.db

	if err := db.Table(taskModel.Task{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": database.Deleted,
		}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
