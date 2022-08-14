package mino

import (
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey string `mapstructure:"accessKey" json:"accessKey" yaml:"accessKey"`
	Secret    string `mapstructure:"secret" json:"secret" yaml:"secret"`
	UseSSL    bool   `mapstructure:"useSSL" json:"useSSL" yaml:"useSSL"`
}

func NewMinio(config *MinioConfig) (*minio.Client, error) {
	option := &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.Secret, ""),
		Secure: config.UseSSL,
	}
	client, err := minio.New(config.Endpoint, option)
	if err != nil {
		fmt.Println("connect minio fail", err)
		return nil, errors.New("connect minio fail")
	}
	return client, nil
}
