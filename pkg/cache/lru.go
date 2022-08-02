package cache

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type LruCache struct {
	Capacity int
	values   *list.List
	cacheMap map[string]*list.Element
	lock     sync.RWMutex
}

func NewLruList(cap int) *LruCache {
	return &LruCache{
		Capacity: cap,
		values:   list.New(),
		cacheMap: make(map[string]*list.Element, cap),
	}
}

func (l *LruCache) Set(k string, v interface{}, expiration time.Duration) error {
	fmt.Println("lru  set")
	l.lock.Lock()
	if element, ok := l.cacheMap[k]; ok {
		l.values.MoveToFront(element)
		element.Value = v
	} else {
		p := l.values.PushFront(v)
		l.cacheMap[k] = p
	}
	l.lock.Unlock()
	if l.values.Len() > l.Capacity {
		l.DelPolicy()
	}
	return nil
}

func (l *LruCache) Get(k string) (interface{}, bool) {
	fmt.Println("lru  get")
	l.lock.RLock()
	defer l.lock.RUnlock()
	element, ok := l.cacheMap[k]
	if ok {
		l.values.MoveToFront(element)
		return element.Value, true
	} else {
		return nil, false
	}
}

func (l *LruCache) Del(k string) bool {
	if l.cacheMap == nil {
		return false
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	if element, ok := l.cacheMap[k]; ok {
		l.values.Remove(element)
		delete(l.cacheMap, k)
		return true
	}
	return false
}

func (l *LruCache) DelPolicy() {
	l.lock.Lock()
	defer l.lock.Unlock()
	var key string
	for k, v := range l.cacheMap {
		if v == l.values.Back() {
			key = k
			break
		}
	}
	l.values.Remove(l.values.Back())
	delete(l.cacheMap, key)
}

func (l *LruCache) Size() int {
	return l.values.Len()
}

func (l *LruCache) List() []interface{} {
	var data []interface{}
	for i := l.values.Front(); i != nil; i = i.Next() {
		data = append(data, i.Value)
	}
	return data
}

func (l *LruCache) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.values = list.New()
	l.cacheMap = make(map[string]*list.Element, l.Capacity)
}
