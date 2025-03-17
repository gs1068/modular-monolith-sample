package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gs1068/modular-monolith-sample/internal/service"
)

type MicropostController struct {
	microsoftService service.MicropostService
}

func NewMicropostController(microsoftService service.MicropostService) *MicropostController {
	return &MicropostController{
		microsoftService: microsoftService,
	}
}

func (mc *MicropostController) GetMicropostByUserID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resErr(c, err, http.StatusBadRequest)
		return
	}

	res, err := mc.microsoftService.GetMicropostByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resData(c, http.StatusOK, res)
}
