package resp

import (
	"net/http"
	"projectName/internal/model"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := response{Code: 0, Message: "success", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, httpCode, code int, message string, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	if errMsg, ok := model.ErrMsgMap[code]; ok && message == "" {
		message = errMsg
	}
	resp := response{Code: code, Message: message, Data: data}
	ctx.JSON(httpCode, resp)
}
