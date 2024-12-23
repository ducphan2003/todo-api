package grpctask

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"todo-api/common"
	"todo-api/middleware"
	"todo-api/modules/task/taskbiz"
	"todo-api/modules/task/taskmodel"
	"todo-api/modules/task/taskstorage"
	"todo-api/proto/taskpb"

	"todo-api/database"
)

type GRPCTaskServer struct {
	taskpb.UnimplementedTodoServiceServer
}

func NewGRPCTaskServer() *GRPCTaskServer {
	return &GRPCTaskServer{}
}

func (s *GRPCTaskServer) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	// md, ok := common.GetMetadataFromContext(ctx)
	// if !ok {
	// 	return nil, common.ErrInvalidRequest(nil)
	// }
	// fmt.Println(md)

	// token := md["Authorization"]

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, common.ErrNoPermission(nil)
	}

	// Lấy token từ metadata
	token := ""
	if val, ok := md["authorization"]; ok {
		token = val[0]
	} else {
		return nil, common.ErrNoPermission(nil)
	}
	fmt.Println(token)
	if token == "" {
		return nil, common.ErrInvalidRequest(nil)
	}

	err := middleware.RequireAuth(ctx, token)
	if err != nil {
		return nil, common.ErrNoPermission(err)
	}

	store := taskstorage.NewSQlStore(database.DB)
	biz := taskbiz.NewCreateBiz(store)

	task := taskmodel.TaskCreate{
		Title:       req.GetTitle(),
		UserID:      uint(req.GetUserId()),
		Description: req.GetDescription(),
		Progress:    req.GetProgress(),
		Priority:    req.GetPriority(),
	}

	if err := biz.CreateTask(ctx, &task); err != nil {
		return nil, common.ErrCannotCreateEntity(taskmodel.Task{}.TableName(), err)
	}

	return &taskpb.CreateTaskResponse{
		Id:          int64(task.ID),
		Title:       task.Title,
		Description: task.Description,
	}, nil
}

func (s *GRPCTaskServer) ReadTask(ctx context.Context, req *taskpb.ReadTaskRequest) (*taskpb.ReadTasksResponse, error) {
	store := taskstorage.NewSQlStore(database.DB)
	biz := taskbiz.NewFindBiz(store)

	taskFilter := taskmodel.TaskFilter{
		UserID:      uint(req.GetUserId()),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Progress:    req.GetProgress(),
		Priority:    req.GetPriority(),
	}

	tasks, err := biz.ListTasksByConditions(ctx, map[string]interface{}{}, &taskFilter)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(taskmodel.Task{}.TableName(), err)
	}

	var tasksRes []*taskpb.ReadTaskResponse
	for _, task := range tasks {
		taskRes := &taskpb.ReadTaskResponse{
			Id:          int64(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Progress:    task.Progress,
			Priority:    task.Priority,
		}
		tasksRes = append(tasksRes, taskRes)
	}
	return &taskpb.ReadTasksResponse{
		Tasks: tasksRes,
	}, nil
}

func (s *GRPCTaskServer) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.ReadTaskResponse, error) {
	store := taskstorage.NewSQlStore(database.DB)
	biz := taskbiz.NewUpdateBiz(store)

	taskData := taskmodel.TaskUpdate{
		UserID:      uint(req.GetUserId()),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Progress:    req.GetProgress(),
		Priority:    req.GetPriority(),
	}

	err := biz.UpdateTask(ctx, uint(req.GetId()), &taskData)
	if err != nil {
		return nil, common.ErrCannotCreateEntity(taskmodel.Task{}.TableName(), err)
	}

	return &taskpb.ReadTaskResponse{
		Id:          int64(taskData.ID),
		Title:       taskData.Title,
		Description: taskData.Description,
		Progress:    taskData.Progress,
		Priority:    taskData.Priority,
	}, nil
}

func (s *GRPCTaskServer) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.EmptyResponse, error) {
	fmt.Println(req)
	store := taskstorage.NewSQlStore(database.DB)
	biz := taskbiz.NewDeleteBiz(store)

	err := biz.DeleteTask(ctx, uint(req.GetId()))
	if err != nil {
		return nil, common.ErrCannotCreateEntity(taskmodel.Task{}.TableName(), err)
	}

	return &taskpb.EmptyResponse{}, nil
}
