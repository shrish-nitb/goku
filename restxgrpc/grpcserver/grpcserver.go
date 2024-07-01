package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "restxgrpc/gen/protos/todopb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

type TodoList map[string]*pb.Task

func NewTodos() *TodoList {
	var init = make(TodoList)
	return &init
}

type server struct {
	List *TodoList
	pb.UnimplementedTodoServiceServer
}

func (s *server) CreateTodo(ctx context.Context, req *pb.TodoMessageRequest) (*pb.TodoListResponse, error) {
	List := *(s.List)
	TodoMessage := req
	fmt.Println(req)
	if strings.TrimSpace(TodoMessage.Task.Value) == "" ||
		strings.TrimSpace(TodoMessage.Id) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request payload")
	}

	if _, exist := List[TodoMessage.Id]; exist {
		return nil, status.Errorf(codes.AlreadyExists, "Resource already exist")
	}

	List[TodoMessage.Id] = &pb.Task{Value: TodoMessage.Task.Value}

	return &pb.TodoListResponse{List: List}, nil
}

func (s *server) ReadTodo(ctx context.Context, req *pb.Empty) (*pb.TodoListResponse, error) {
	List := *(s.List)
	return &pb.TodoListResponse{List: List}, nil
}

func (s *server) UpdateTodo(ctx context.Context, req *pb.TodoMessageRequest) (*pb.TodoListResponse, error) {
	List := *(s.List)
	TodoMessage := req
	if strings.TrimSpace(TodoMessage.Task.Value) == "" || strings.TrimSpace(TodoMessage.Id) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request payload")
	}

	if _, exist := List[TodoMessage.Id]; !exist {
		return nil, status.Errorf(codes.NotFound, "Resource not exist")
	}

	List[TodoMessage.Id] = &pb.Task{Value: TodoMessage.Task.Value}

	return &pb.TodoListResponse{List: List}, nil
}

func (s *server) DeleteTodo(ctx context.Context, req *pb.TodoMessageRequest) (*pb.TodoListResponse, error) {
	List := *(s.List)
	TodoMessage := req
	if strings.TrimSpace(TodoMessage.Id) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request payload")
	}

	if _, exist := List[TodoMessage.Id]; !exist {
		return nil, status.Errorf(codes.NotFound, "Resource not exist")
	}

	delete(List, TodoMessage.Id)

	return &pb.TodoListResponse{List: List}, nil
}

func Run() {
	todoList := NewTodos()

	lis, err := net.Listen("tcp", "127.0.0.1:5005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{
		List: todoList,
	})

	log.Printf("gRPC server listening at %v", lis.Addr())
	go s.Serve(lis)
}
