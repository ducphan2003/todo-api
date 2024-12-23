package userbiz

import (
	"context"
	"fmt"
	"todo-api/common"
	userModel "todo-api/modules/user/usermodel"
)

type CreateStorage interface {
	CreateUser(ctx context.Context, data *userModel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type createBiz struct {
	createStorage CreateStorage
	hasher        Hasher
}

func NewCreateBiz(createStorage CreateStorage, hasher Hasher) *createBiz {
	return &createBiz{
		createStorage: createStorage,
		hasher:        hasher,
	}
}

func (biz createBiz) CreateUser(ctx context.Context, data *userModel.UserCreate) error {
	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	if err := biz.createStorage.CreateUser(ctx, data); err != nil {
		fmt.Println(err)
		return common.ErrCannotCreateEntity(userModel.User{}.TableName(), err)
	}

	return nil
}
