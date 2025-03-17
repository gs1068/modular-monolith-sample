package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gs1068/modular-monolith-sample/internal/adapter/controller"
)

func setupUserRoutes(api *gin.RouterGroup, uc *controller.UserController) {
	users := api.Group("/users")
	users.GET("/:id", uc.GetUserByID)
}
