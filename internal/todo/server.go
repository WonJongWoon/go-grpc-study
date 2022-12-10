package todo

import (
	"context"
	"errors"
	pb "go-grpc-study/protos/v1/todo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.TodoServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) GetTodo(ctx context.Context, request *pb.GetTodoRequest) (*pb.TodoResponse, error) {
	var todo Todo
	if err := s.db.First(&todo, uint(request.TodoId)).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "todo id '%d' does not exist", request.TodoId)
	}

	return &pb.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
		Status:      pb.TodoStatus(pb.TodoStatus_value[string(todo.Status)]),
		Author:      todo.Author,
		CreatedAt:   todo.CreatedAt.String(),
		ModifiedAt:  todo.UpdatedAt.String(),
	}, nil
}

func (s *Server) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.TodoResponse, error) {
	newTodo := New(request.Title, request.Description, "woony")
	if err := s.db.Create(newTodo).Error; err != nil {
		return nil, status.Error(codes.Internal, "todo failed to create")
	}

	return &pb.TodoResponse{
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Status:      pb.TodoStatus(pb.TodoStatus_value[string(newTodo.Status)]),
		Author:      newTodo.Author,
		CreatedAt:   newTodo.CreatedAt.String(),
		ModifiedAt:  newTodo.UpdatedAt.String(),
	}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, request *pb.UpdateTodoRequest) (*pb.TodoResponse, error) {
	var todo Todo
	if err := s.db.First(&todo, uint(request.TodoId)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "todo id '%d' does not exist", request.TodoId)
		}
		return nil, status.Errorf(codes.Internal, "todo failed to get")
	}

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		todo.Description = request.Description
	}

	todo.Status = Status(pb.TodoStatus_name[int32(request.Status)])

	if err := s.db.Save(&todo).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "todo id '%d' failed to update", request.TodoId)
	}

	return &pb.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
		Status:      pb.TodoStatus(pb.TodoStatus_value[string(todo.Status)]),
		Author:      todo.Author,
		CreatedAt:   todo.CreatedAt.String(),
		ModifiedAt:  todo.UpdatedAt.String(),
	}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, request *pb.DeleteTodoRequest) (*pb.EmptyResponse, error) {
	if err := s.db.First(&Todo{}, uint(request.TodoId)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "todo id '%d' does not exist", request.TodoId)
		}
		return nil, status.Errorf(codes.Internal, "todo id '%d' failed to delete", request.TodoId)
	}

	if err := s.db.Delete(&Todo{}, uint(request.TodoId)).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "todo id '%d' failed to delete", request.TodoId)
	}

	return &pb.EmptyResponse{}, nil
}
