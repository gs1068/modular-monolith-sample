package service

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/internal/adapter/dto"

	userpb "github.com/gs1068/modular-monolith-sample/gen/user/v1"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (*dto.UserOutput, error)
}

type userService struct {
	client userpb.UserServiceServer
}

func NewUserService(client userpb.UserServiceServer) UserService {
	return &userService{
		client: client,
	}
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*dto.UserOutput, error) {
	res, err := s.client.GetUserByID(ctx, &userpb.GetUserByIDRequest{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}
	return dto.ConvertUserResponse(res.User), nil
}
