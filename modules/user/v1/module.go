package module

import (
	"time"

	userpb "github.com/gs1068/modular-monolith-sample/gen/user/v1"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/handler"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/infrastructure/mysql"
	"github.com/gs1068/modular-monolith-sample/modules/user/internal/usecase"
)

func NewUserModule() (userpb.UserServiceServer, error) {
	userDB, err := mysql.NewDB("root:@tcp(localhost:3306)/user_db")
	if err != nil {
		return nil, err
	}
	userDB.SetMaxOpenConns(0)
	userDB.SetMaxIdleConns(100)
	userDB.SetConnMaxLifetime(10 * time.Second)

	repo := mysql.NewUserRepository(userDB)
	uc := usecase.NewUserUsecase(repo)
	return handler.NewUserServiceServer(uc), nil
}
