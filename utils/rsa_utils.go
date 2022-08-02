package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"memoirs/global"
)

// 生成RSA密钥对
func GeneratorRSAKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	key1 := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: key1,
	}
	prvKey := pem.EncodeToMemory(block)

	publicKey := privateKey.PublicKey
	key2, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: key2,
	}
	pubKey := pem.EncodeToMemory(block)
	return prvKey, pubKey, err
}

func RsaEncrypt(data, keyByte []byte) ([]byte, error) {
	block, _ := pem.Decode(keyByte)
	if block == nil {
		global.Log.Error("RSA公钥错误")
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	return ciphertext, err
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		global.Log.Error("RSA私钥错误")
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	return data, err
}
