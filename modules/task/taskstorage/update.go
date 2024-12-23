package taskstorage

import (
	"context"
	"todo-api/common"
	taskModel "todo-api/modules/task/taskmodel"
)

func (s *sqlStore) UpdateTask(ctx context.Context,
	id uint,
	data *taskModel.TaskUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
