package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type any = interface{}

var (
	NULL_DATA  = map[string]any{}
	STATUS_OK  = 200
	SERVER_ERR = 500
)

type Response struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 响应数据
}

func resultJson(ctx *gin.Context, status, code int, message string, data any) {
	ctx.JSON(status, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Ok(ctx *gin.Context) {
	resultJson(ctx, http.StatusOK, STATUS_OK, "success", NULL_DATA)
}

func OkWithData(ctx *gin.Context, data any) {
	resultJson(ctx, http.StatusOK, STATUS_OK, "success", data)
}

func OkWithDetail(ctx *gin.Context, message string, data any) {
	resultJson(ctx, http.StatusOK, STATUS_OK, message, data)
}

func FailWithMessage(ctx *gin.Context, message string) {
	resultJson(ctx, http.StatusInternalServerError, SERVER_ERR, message, NULL_DATA)
}

func FailWithDetail(ctx *gin.Context, status int, message string) {
	resultJson(ctx, status, status, message, NULL_DATA)
}
