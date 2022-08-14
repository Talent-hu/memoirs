package mino

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"testing"
)

var client *minio.Client

func TestNewMinio(t *testing.T) {
	config := &MinioConfig{
		Endpoint:  "101.132.251.60:9000",
		AccessKey: "minio",
		Secret:    "minio123321",
		UseSSL:    false,
	}
	var err error
	client, err = NewMinio(config)
	if err != nil {
		t.Fatal("连接minio失败！", err)
	}
	createBucket("images")
	//listBucket()
	//FileUpload()
}

func createBucket(bucketName string) {
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{ObjectLocking: false})
	if err != nil {
		log.Println("创建bucket错误：", err)
		exists, _ := client.BucketExists(context.Background(), bucketName)
		if exists {
			log.Printf("bucket:%s已经存在\n", bucketName)
		}
	} else {
		log.Printf("successfully create %s\n", bucketName)
	}
}

func listBucket() {
	buckets, _ := client.ListBuckets(context.Background())
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func FileUpload() {
	bucketName := "memoirs"
	objectName := "aaa.png"
	filePath := "./下载.png"
	contentType := "image/png"
	object, err := client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Printf("上传失败！err:%s", err)
	}
	fmt.Println(object)
}
