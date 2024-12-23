package userbiz

import (
	"context"
	// "encoding/json"
	// "os"
	"todo-api/common"
	"todo-api/component/tokenprovider"
	// "todo-api/component/tokenprovider/jwt"
	userModel "todo-api/modules/user/usermodel"
)

type LoginStorage interface {
	FindOneUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*userModel.User, error)
}

type TokenConfig interface {
	GetAccessTokenExp() int
	GetAccessTokenPublicKey() string
	GetAccessTokenPrivateKey() string
}

type loginBusiness struct {
	loginStorage  LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tokenConfig   TokenConfig
}

func NewLoginBiz(storage LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	tokenConfig TokenConfig,
) *loginBusiness {
	return &loginBusiness{
		loginStorage:  storage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tokenConfig:   tokenConfig,
	}
}

func (biz *loginBusiness) Login(ctx context.Context,
	data *userModel.UserLogin,
) (*tokenprovider.Token, error) {

	user, err := biz.loginStorage.FindOneUser(ctx, map[string]interface{}{"name": data.Name})
	if err != nil {
		return nil, userModel.AccountNotFound
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, userModel.ErrUserNameOrPasswordInvalid
	}
	if (user != nil) && (!user.DeletedAt.Time.IsZero()) {
		return nil, userModel.AccountDeleted
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tokenConfig.GetAccessTokenExp(), biz.tokenConfig.GetAccessTokenPrivateKey())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
