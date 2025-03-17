package handler

import (
	"context"

	micropostpb "github.com/gs1068/modular-monolith-sample/gen/micropost/v1"
	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/usecase"
)

type MicropostServiceServer struct {
	micropostUsecase *usecase.MicropostUsecase
	micropostpb.UnimplementedMicropostServiceServer
}

func NewMicropostServiceServer(micopostUsecase *usecase.MicropostUsecase) *MicropostServiceServer {
	return &MicropostServiceServer{
		micropostUsecase: micopostUsecase,
	}
}

func (s *MicropostServiceServer) GetMicropostByUserID(ctx context.Context, req *micropostpb.GetMicropostByUserIDRequest) (*micropostpb.GetMicropostByUserIDResponse, error) {
	ms, err := s.micropostUsecase.GetMicropostByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	var res []*micropostpb.Micropost
	for _, m := range ms {
		res = append(res, &micropostpb.Micropost{
			Id:      m.ID,
			UserId:  m.UserID,
			Content: m.Content,
		})
	}

	return &micropostpb.GetMicropostByUserIDResponse{
		Microposts: res,
	}, nil
}
