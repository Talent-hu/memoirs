package cache

import (
	"errors"
	"fmt"
	"memoirs/global"
	"time"
)

var (
	REDIS = "redis"
	LRU   = "lru"
)

var CacheManager *CacheManage

func InitCache() {
	cacheMode := global.Config.CacheMode
	cacheSize := global.Config.CacheSize
	CacheManager = NewCacheMange(cacheMode, cacheSize)
}

type CacheManage struct {
	cache Cache
}

func NewCacheMange(cacheMode string, cacheSize int) *CacheManage {
	fmt.Println("start load cache manager")
	if cacheMode == LRU {
		if cacheSize <= 0 {
			panic(errors.New("must be set cache size"))
		}
		return &CacheManage{
			cache: NewLruList(cacheSize),
		}
	} else if cacheMode == REDIS {
		return &CacheManage{
			cache: NewRedisCache(),
		}
	}
	return nil
}

func (this *CacheManage) Set(key string, val interface{}, expiration time.Duration) (err error) {
	return this.cache.Set(key, val, expiration)
}

func (this *CacheManage) Get(key string) (val interface{}, b bool) {
	return this.cache.Get(key)
}

func (this *CacheManage) Del(key string) (b bool) {
	return this.cache.Del(key)
}
