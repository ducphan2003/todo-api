package grpcuser

import (
	"context"
	"todo-api/common"
	"todo-api/component/hasher"
	"todo-api/component/tokenprovider"
	"todo-api/component/tokenprovider/jwt"
	"todo-api/modules/user/userbiz"
	"todo-api/modules/user/usermodel"
	"todo-api/modules/user/userstorage"
	"todo-api/proto/userpb"

	"todo-api/database"
)

type GRPCUserServer struct {
	userpb.UnimplementedTodoServiceServer
}

func NewGRPCUserServer() *GRPCUserServer {
	return &GRPCUserServer{}
}

func (s *GRPCUserServer) SignUp(ctx context.Context, req *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {
	store := userstorage.NewSQlStore(database.DB)
	md5 := hasher.NewMd5Hash()
	biz := userbiz.NewCreateBiz(store, md5)

	user := usermodel.UserCreate{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	if err := biz.CreateUser(ctx, &user); err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.User{}.TableName(), err)
	}

	return &userpb.SignUpResponse{
		Name:     user.Name,
		Password: "",
	}, nil
}

func (s *GRPCUserServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	store := userstorage.NewSQlStore(database.DB)
	md5 := hasher.NewMd5Hash()
	tokenConfig := tokenprovider.NewTokenConfig()
	tokenProvider := jwt.NewTokenJWTProvider(tokenConfig.GetAccessTokenPrivateKey(), tokenConfig.GetAccessTokenPublicKey())
	biz := userbiz.NewLoginBiz(store, tokenProvider, md5, tokenConfig)

	user := usermodel.UserLogin{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	token, err := biz.Login(ctx, &user)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.User{}.TableName(), err)
	}

	return &userpb.LoginResponse{
		Token: token.Token,
	}, nil
}
