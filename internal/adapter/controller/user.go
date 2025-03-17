package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gs1068/modular-monolith-sample/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (h *UserController) GetUserByID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resErr(c, err, http.StatusBadRequest)
		return
	}

	res, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		resErr(c, err, http.StatusInternalServerError)
		return
	}

	resData(c, http.StatusOK, res)
}
