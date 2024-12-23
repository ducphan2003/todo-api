package taskstorage

import (
	"context"
	"gorm.io/gorm"
	"todo-api/common"

	taskModel "todo-api/modules/task/taskmodel"
)

func (s *sqlStore) FindOneTask(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*taskModel.Task, error) {
	db := s.db.Table(taskModel.Task{}.TableName())
	for i := range moreInfo {
		db.Preload(moreInfo[i])
	}
	var task taskModel.Task
	if err := db.Where(conditions).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &task, nil
}

func (s *sqlStore) ListTasksByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	filter *taskModel.TaskFilter,
	moreKeys ...string,
) ([]taskModel.Task, error) {
	filter.Status = "active"
	var result []taskModel.Task
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Where(conditions)
	if v := filter; v != nil {
		if v.UserID > 0 {
			db.Where("user_id = ?", v.UserID)
		}
		if v.Status != "" {
			db.Where("status = ?", v.Status)
		}
		if v.Title != "" {
			db.Where("title = ?", v.Title)
		}
		if v.Priority != "" {
			db.Where("priority = ?", v.Priority)
		}
		if v.Progress != "" {
			db.Where("progress = ?", v.Progress)
		}
	}
	db.Model(&result)
	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
