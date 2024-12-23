package routes

import (
	"todo-api/modules/task/transportation/grpctask"
	"todo-api/modules/user/transportation/grpcuser"
	"todo-api/proto/taskpb"
	"todo-api/proto/userpb"

	"google.golang.org/grpc"
)

func RegisterGRPCRoutes(s *grpc.Server) {
	userpb.RegisterTodoServiceServer(s, grpcuser.NewGRPCUserServer())
	taskpb.RegisterTodoServiceServer(s, grpctask.NewGRPCTaskServer())
}
