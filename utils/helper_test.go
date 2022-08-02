package utils

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	fmt.Println(uuid)
}

func TestGenerateMD5(t *testing.T) {
	password := "123456"
	encPwd := GenerateMD5(password)
	fmt.Println(encPwd)
}

func TestGeneratorRSAKey(t *testing.T) {
	privateKey, publicKey, _ := GeneratorRSAKey()
	fmt.Println(string(privateKey))
	fmt.Println(string(publicKey))
	password := "123456"
	encPwd, _ := RsaEncrypt([]byte(password), publicKey)
	fmt.Println(string(encPwd))
	pwd, _ := RsaDecrypt(encPwd, privateKey)
	fmt.Println(string(pwd))

}
