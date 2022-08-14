package file

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/pkg/response"
)

type ImageApi struct{}

func (image *ImageApi) ImageUpload(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.Log.Error("接收文件失败！", zap.Error(err))
		response.FailWithMessage(ctx, "接收文件失败！")
		return
	}
	file, err := fileService.UploadFile(header, ImageBucket)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, file)
}
