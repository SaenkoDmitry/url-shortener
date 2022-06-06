package cache

import (
	"errors"

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
	Set(key string, value interface{}) error
}

type inMemoryCache struct {
	l *lru.Cache
}

func (im *inMemoryCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (im *inMemoryCache) Set(key string, value interface{}) error {
	return errors.New("TODO")
}
