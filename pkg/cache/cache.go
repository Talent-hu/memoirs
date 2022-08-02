package cache

import "time"

type Cache interface {
	Set(key string, val interface{}, expiration time.Duration) error
	Get(key string) (interface{}, bool)
	Del(key string) bool
}
