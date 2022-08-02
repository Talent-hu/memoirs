package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	u := uuid.New()
	return u.String()
}

func GenerateMD5(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum([]byte(pwd)))
}
