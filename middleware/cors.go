package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method                                                                                                           //请求方法
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))                                                                     // 指明哪些请求源被允许访问资源，值可以为 "*"（允许访问所有域），"null"，或者单个源地址。
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")                              // 对于预请求来说，指明了哪些头信息可以用于实际的请求中。
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                                                       // 对于预请求来说，哪些请求方式可以用于实际的请求。
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") // 对于预请求来说，指明哪些头信息可以安全的暴露给 CORS API 规范的 API
		ctx.Header("Access-Control-Allow-Credentials", "true")                                                                                 // 指明当请求中省略 creadentials 标识时响应是否暴露。对于预请求来说，它表明实际的请求中可以包含用户凭证。
		// c.Header("Access-Control-Max-Age", "172800")                                                                                         // 缓存请求信息 单位为秒
		// c.Set("content-type", "application/json")                                                                                            // 设置返回格式是json
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}
