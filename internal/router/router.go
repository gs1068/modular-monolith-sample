package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gs1068/modular-monolith-sample/internal/adapter/controller"
)

func SetupRoutes(
	engine *gin.Engine,
	uc *controller.UserController,
	mc *controller.MicropostController,
) {
	api := engine.Group("/api/v1")
	setupUserRoutes(api, uc)
	setupMicropostRoutes(api, mc)
}
