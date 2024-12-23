package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"todo-api/common"
	"todo-api/component/tokenprovider"
	"todo-api/component/tokenprovider/jwt"
	"todo-api/database"
	userBiz "todo-api/modules/user/userbiz"
	// "todo-api/modules/user/usermodel"
	userStorage "todo-api/modules/user/userstorage"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthAHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(ctx context.Context, token string) error {
	userStore := userStorage.NewSQlStore(database.DB)

	biz := userBiz.NewFindBiz(userStore)
	tokenConfig := tokenprovider.NewTokenConfig()
	tokenProvider := jwt.NewTokenJWTProvider(tokenConfig.GetAccessTokenPrivateKey(), tokenConfig.GetAccessTokenPublicKey())
	payload, err := tokenProvider.Validate(token, tokenConfig.GetAccessTokenPublicKey())
	if err != nil {
		return err
	}
	user, err := biz.FindUser(ctx, map[string]interface{}{"id": payload.UserId})
	if err != nil {
		return err
	}
	if user.Status == database.Deleted {
		return common.ErrNoPermission(errors.New("user has been deleted"))
	}
	return nil
}
