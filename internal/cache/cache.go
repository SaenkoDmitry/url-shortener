package cache

import (
	lru "github.com/hashicorp/golang-lru"
)

func NewCache(size int) (MyCache, error) {
	l, err := lru.New(size)
	if err != nil {
		return nil, err
	}

	return &inMemoryCache{
		l: l,
	}, nil
}

type MyCache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type inMemoryCache struct {
	l *lru.Cache
}

func (im *inMemoryCache) Get(key string) (interface{}, bool) {
	return im.l.Get(key)
}

func (im *inMemoryCache) Set(key string, value interface{}) {
	im.l.Add(key, value)
}
