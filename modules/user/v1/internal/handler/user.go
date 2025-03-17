package handler

import (
	"context"

	userpb "github.com/gs1068/modular-monolith-sample/gen/user/v1"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	userUsecase *usecase.UserUsecase
	userpb.UnimplementedUserServiceServer
}

func NewUserServiceServer(userUsecase *usecase.UserUsecase) *UserServiceServer {
	return &UserServiceServer{
		userUsecase: userUsecase,
	}
}

func (s *UserServiceServer) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := s.userUsecase.GetUserByID(req.GetUserId(), ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}

	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %d", req.GetUserId())
	}

	return &userpb.GetUserByIDResponse{
		User: &userpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}
