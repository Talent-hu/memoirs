package cache

import (
	"fmt"
	"testing"
)

func TestCacheManage_Get(t *testing.T) {
	cacheMode := "lru"
	cacheSize := 1024
	mange := NewCacheMange(cacheMode, cacheSize)
	err := mange.Set("name", "alex")
	if err != nil {
		fmt.Println(err)
	}
	name, b := mange.Get("name")
	if b {
		fmt.Println(name)
	}

}
