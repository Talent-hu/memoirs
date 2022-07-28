package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"memoirs/global"
	"net/http"
	"strconv"
	"strings"
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
		blw := &BodyLogWrite{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		jsonStr := fmt.Sprintf("%s\n", strings.ReplaceAll(blw.body.String(), "\\\"", "\""))
		fmt.Println("响应数据为：", jsonStr)
		global.Log.Info("read body from response:",
			zap.String("url", ctx.Request.URL.String()),
			zap.String("status", strconv.Itoa(ctx.Writer.Status())),
			zap.String("responseBody", blw.body.String()))
	}
}

type BodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWrite) Write(b []byte) (int, error) {
	bf := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(bf)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(b)
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w BodyLogWrite) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
