package usecase

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/domain/model"
	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/domain/repository"
)

type MicropostUsecase struct {
	micropostRepo repository.MicropostRepository
}

func NewMicropostUsecase(micropostRepo repository.MicropostRepository) *MicropostUsecase {
	return &MicropostUsecase{
		micropostRepo: micropostRepo,
	}
}

func (u *MicropostUsecase) GetMicropostByUserID(ctx context.Context, userID int64) ([]*model.Micropost, error) {
	res, err := u.micropostRepo.GetMicropostByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
