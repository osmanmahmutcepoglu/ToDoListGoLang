package controllers

import (
	"github.com/gin-gonic/gin"
)

type Attachment struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ControllerBase interface {
	Data(c *gin.Context, code int, data interface{}, message string)
	Error(c *gin.Context, code int, errorMessasge error)
}

type controllerBase struct{}

func NewBaseController() ControllerBase {
	return &controllerBase{}
}

func (_basecontroller *controllerBase) Data(c *gin.Context, code int, data interface{}, message string) {
	attach := &Attachment{
		Code:    code,
		Data:    data,
		Message: message,
	}

	c.JSON(code, attach)
}

func (_basecontroller *controllerBase) Error(c *gin.Context, code int, errorMessasge error) {
	attach := &Attachment{
		Code:    code,
		Data:    nil,
		Message: errorMessasge.Error(),
	}

	c.AbortWithStatusJSON(code, attach)
}
