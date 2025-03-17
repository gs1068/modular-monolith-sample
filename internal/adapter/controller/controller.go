package controller

import (
	"github.com/gin-gonic/gin"
)

func resErr(c *gin.Context, err error, status int) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func resData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
