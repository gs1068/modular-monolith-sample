package service

import (
	"context"

	micropostpb "github.com/gs1068/modular-monolith-sample/gen/micropost/v1"
	userpb "github.com/gs1068/modular-monolith-sample/gen/user/v1"
	"github.com/gs1068/modular-monolith-sample/internal/adapter/dto"
)

type MicropostService interface {
	GetMicropostByUserID(ctx context.Context, userID int64) (*dto.GetMicropostByUserIDOutput, error)
}

type micropostService struct {
	mc micropostpb.MicropostServiceServer
	uc userpb.UserServiceServer
}

func NewMicropostService(
	mc micropostpb.MicropostServiceServer,
	uc userpb.UserServiceServer,
) MicropostService {
	return &micropostService{
		mc: mc,
		uc: uc,
	}
}

func (s *micropostService) GetMicropostByUserID(ctx context.Context, userID int64) (*dto.GetMicropostByUserIDOutput, error) {
	ms, err := s.mc.GetMicropostByUserID(ctx, &micropostpb.GetMicropostByUserIDRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	u, err := s.uc.GetUserByID(ctx, &userpb.GetUserByIDRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return dto.ConvertGetMicropostByUserIDResponse(ms.Microposts, u.User), nil
}
