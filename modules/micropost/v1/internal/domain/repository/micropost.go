package repository

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/domain/model"
)

type MicropostRepository interface {
	GetMicropostByUserID(ctx context.Context, userID int64) ([]*model.Micropost, error)
}
