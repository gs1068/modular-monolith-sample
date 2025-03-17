package usecase

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/user/internal/domain/model"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/domain/repository"
)

type UserUsecase struct {
	useRepo repository.UserRepository
}

func NewUserUsecase(useRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		useRepo: useRepo,
	}
}

func (u *UserUsecase) GetUserByID(id int64, ctx context.Context) (*model.User, error) {
	res, err := u.useRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
