package repository

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/user/internal/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}
