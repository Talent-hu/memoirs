package file

import "memoirs/service"

// minio存图片的桶
var ImageBucket = "images"

var (
	fileService = service.ServiceGroupApp.FileService
)
