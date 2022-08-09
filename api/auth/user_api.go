package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"memoirs/global"
	"memoirs/model/vo"
	"memoirs/pkg/response"
	"memoirs/utils"
	"memoirs/validate"
)

type UserApi struct{}

// @Tag UserApi
// @Summary 用户登录
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param data body vo.LoginRequest false "查询参数"
// @Success 200 {object} response.Response{data=vo.LoginResponse,message=string} "返回用户token"
// @Router /user/login [post]
func (this *UserApi) Login(ctx *gin.Context) {
	var loginReq vo.LoginRequest
	_ = ctx.ShouldBindJSON(&loginReq)
	if err := validate.Verify(loginReq, validate.LoginVerify); err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	resp, err := authService.Login(loginReq)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	} else {
		response.OkWithData(ctx, resp)
	}
}

// @Tag UserApi
// @Summary 获取RSA公钥
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.GetRsaKeyResponse,message=string} "返回用户token"
// @Router /user/publicKey [get]
func (this *UserApi) PublicKey(ctx *gin.Context) {
	resp, err := authService.RsaPublicSecret()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.OkWithData(ctx, resp)
}

func (this *UserApi) Logout(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	identity := claims.UserInfo.Identity
	err = global.Redis.Del(context.Background(), identity).Err()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.Ok(ctx)
}

// @Tag UserApi
// @Summary 获取用户信息
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} response.Response{data=vo.GetUserInfoResponse,message=string} "返回用户token"
// @Router /user/info [get]
func (this *UserApi) QueryUserInfo(ctx *gin.Context) {
	userId := utils.GetUserID(ctx)
	resp, err := userService.QueryUserInfo(userId)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.OkWithData(ctx, resp)
}

func (this *UserApi) QueryUserList(ctx *gin.Context) {
	var pageQuery vo.ListQuery
	_ = ctx.ShouldBindJSON(&pageQuery)
	resp, err := userService.QueryUserList(pageQuery)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.OkWithData(ctx, resp)
}
