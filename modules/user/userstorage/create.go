package userstorage

import (
	"context"
	"todo-api/common"
	userModel "todo-api/modules/user/usermodel"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *userModel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(userModel.User{}.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
