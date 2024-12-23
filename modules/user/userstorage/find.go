package userstorage

import (
	"context"
	"gorm.io/gorm"
	"todo-api/common"

	userModel "todo-api/modules/user/usermodel"
)

func (s *sqlStore) FindOneUser(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*userModel.User, error) {
	db := s.db.Table(userModel.User{}.TableName())
	for i := range moreInfo {
		db.Preload(moreInfo[i])
	}
	var user userModel.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
