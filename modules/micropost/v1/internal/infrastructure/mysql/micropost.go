package mysql

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/domain/model"
	"github.com/jmoiron/sqlx"
)

type micropostRepository struct {
	db *sqlx.DB
}

func NewMicropostRepository(db *sqlx.DB) *micropostRepository {
	return &micropostRepository{
		db: db,
	}
}

type resMicropost struct {
	ID      int64  `db:"id"`
	UserID  int64  `db:"user_id"`
	Content string `db:"content"`
}

const (
	getMicropostByUserIDQuery = `SELECT id, content, user_id FROM microposts WHERE user_id = ?`
)

func (r *micropostRepository) GetMicropostByUserID(ctx context.Context, userID int64) ([]*model.Micropost, error) {
	var rm []resMicropost
	if err := r.db.SelectContext(ctx, &rm, getMicropostByUserIDQuery, userID); err != nil {
		return nil, err
	}

	res := make([]*model.Micropost, 0, len(rm))
	for _, m := range rm {
		res = append(res, &model.Micropost{
			ID:      m.ID,
			UserID:  m.UserID,
			Content: m.Content,
		})
	}

	return res, nil
}
