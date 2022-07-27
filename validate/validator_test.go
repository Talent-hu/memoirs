package validate

import (
	"fmt"
	"testing"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func TestVerify(t *testing.T) {
	user := User{
		Username: "admin",
		Password: "asad",
		Age:      -24,
	}

	err := Verify(user, LoginVerify)
	if err != nil {
		fmt.Println(err)
	}
}
