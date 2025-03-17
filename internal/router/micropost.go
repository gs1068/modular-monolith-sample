package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gs1068/modular-monolith-sample/internal/adapter/controller"
)

func setupMicropostRoutes(api *gin.RouterGroup, uc *controller.MicropostController) {
	m := api.Group("/microposts")
	m.GET("/:id", uc.GetMicropostByUserID)
}
