package utils

import (
	"fmt"
	"testing"
)

type User struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Age   uint   `json:"age"`
}

type UserInfo struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Status *bool  `json:"status"`
	Age    uint   `json:"age"`
}

func TestUserInfo(t *testing.T) {
	user := new(UserInfo)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%#v\n", user.Status)
}

func TestBeanUtils(t *testing.T) {
	user := &User{
		Name:  "admin",
		Title: "666",
		Age:   23,
	}
	var userInfo UserInfo
	err := CopyProperties(&userInfo, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userInfo)
}
