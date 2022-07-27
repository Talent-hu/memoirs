package jwts

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserId      uint
	UserName    string
	NickName    string
	Identity    string
}

type UserStdClaims struct {
	jwt.StandardClaims
	BufferTime int64
	UserInfo   *UserClaims
}
