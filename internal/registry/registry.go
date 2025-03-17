package registry

import (
	"github.com/gin-gonic/gin"
	"github.com/gs1068/modular-monolith-sample/internal/adapter/controller"
	"github.com/gs1068/modular-monolith-sample/internal/router"
	"github.com/gs1068/modular-monolith-sample/internal/service"
	micropostModule "github.com/gs1068/modular-monolith-sample/modules/micropost"
	userModule "github.com/gs1068/modular-monolith-sample/modules/user"
)

func NewRegistry(engine *gin.Engine) {
	um, err := userModule.NewUserModule()
	if err != nil {
		panic(err)
	}

	mm, err := micropostModule.NewMicropostModule()
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService(um)
	micropostService := service.NewMicropostService(mm, um)

	userController := controller.NewUserController(userService)
	micropostController := controller.NewMicropostController(micropostService)

	router.SetupRoutes(engine, userController, micropostController)
}
