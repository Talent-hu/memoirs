package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"memoirs/global"
	"net/http"
)

func RecordOptions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		if method == http.MethodPost {
			body, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				global.Log.Error("read body from request error:", zap.Error(err))
			} else {
				global.Log.Info("read body from request:", zap.String("requestBody", string(body)))
				ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		blw := &bodyLogWrite{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		fmt.Println(ctx.Writer.Status())
		global.Log.Info("read body from response:", zap.String("responseBody", blw.body.String()))
	}
}

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
