package file

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/pkg/conf"
	"mime/multipart"
)

type FileService struct{}

func (srv *FileService) UploadFile(header *multipart.FileHeader, bucketName string) (map[string]string, error) {
	err := srv.hasBucket(bucketName)
	if err != nil {
		return nil, err
	}
	filename := header.Filename
	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		return nil, errors.New("只支持image/jpeg、image/png格式图片上传")
	}
	opts := minio.PutObjectOptions{ContentType: contentType}
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = global.Minio.PutObject(context.Background(), bucketName, filename, file, header.Size, opts)
	if err != nil {
		global.Log.Error("文件上传失败", zap.Error(err))
		return nil, errors.New("文件上传失败")
	}
	cfg, _ := conf.NewMinioCfg(global.Viper)
	fileUrl := fmt.Sprintf("http://%s/%s/%s", cfg.Endpoint, bucketName, filename)
	data := map[string]string{}
	data["fileUrl"] = fileUrl
	return data, nil
}

func (srv *FileService) hasBucket(bucketName string) error {
	exists, _ := global.Minio.BucketExists(context.Background(), bucketName)
	if !exists {
		fmt.Printf("文件bucket:%s不存在，请联系管理员修复", bucketName)
		return errors.New("文件bucket不存在，请联系管理员修复")
	}
	return nil
}
