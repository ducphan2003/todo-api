package userbiz

import (
	"context"
	"todo-api/common"
	userModel "todo-api/modules/user/usermodel"
)

type FindStorage interface {
	FindOneUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*userModel.User, error)
}

type findBiz struct {
	findStorage FindStorage
}

func NewFindBiz(findStorage FindStorage) *findBiz {
	return &findBiz{
		findStorage: findStorage,
	}
}

func (biz findBiz) FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	user, err := biz.findStorage.FindOneUser(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(userModel.User{}.TableName(), err)
	}

	return user, nil
}
