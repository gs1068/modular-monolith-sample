package module

import (
	"time"

	micropostpb "github.com/gs1068/modular-monolith-sample/gen/micropost/v1"
	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/handler"
	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/infrastructure/mysql"
	"github.com/gs1068/modular-monolith-sample/modules/micropost/internal/usecase"
)

func NewMicropostModule() (micropostpb.MicropostServiceServer, error) {
	micropostDB, err := mysql.NewDB("root:@tcp(localhost:3306)/micropost_db")
	if err != nil {
		return nil, err
	}
	micropostDB.SetMaxOpenConns(0)
	micropostDB.SetMaxIdleConns(100)
	micropostDB.SetConnMaxLifetime(10 * time.Second)

	repo := mysql.NewMicropostRepository(micropostDB)
	uc := usecase.NewMicropostUsecase(repo)
	return handler.NewMicropostServiceServer(uc), nil
}
