package userstorage

import (
	"context"
	"todo-api/common"
	userModel "todo-api/modules/user/usermodel"
)

func (s *sqlStore) UpdateUser(ctx context.Context,
	id uint,
	data *userModel.UserUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
