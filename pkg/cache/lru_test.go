package cache

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	list := NewLruList(3)
	list.Set("name", "alex")
	list.Set("age", "100")
	list.Set("gender", "男")
	list.Set("nickName", "666")
	name, _ := list.Get("name")
	fmt.Println(name)

	age, _ := list.Get("age")
	fmt.Println(age)

	gender, _ := list.Get("gender")
	fmt.Println(gender)

	nickName, _ := list.Get("nickName")
	fmt.Println(nickName)
}
