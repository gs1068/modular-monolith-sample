package mysql

import (
	"context"

	"github.com/gs1068/modular-monolith-sample/modules/user/internal/domain/model"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

const (
	getUserByIDQuery = `SELECT id, email, name FROM users WHERE id = ?`
)

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, getUserByIDQuery, id); err != nil {
		return nil, err
	}
	return &user, nil
}
