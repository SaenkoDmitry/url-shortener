package service

import (
	"fmt"
	"time"
	"url-shortener/internal/cache"
	"url-shortener/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	GeneratedURLLen = 7 // need to calculate this number by combination requests load and durability of the service
)

type URLService struct {
	cacheStore    cache.MyCache
	shortURLStore cache.MyCache
}

func NewURLService(cacheStore, shortURLStore cache.MyCache) *URLService {
	return &URLService{
		cacheStore:    cacheStore,
		shortURLStore: shortURLStore,
	}
}

const (
	CreateURIErr     = "cannot create uri: %v"
	SuchURLNotExists = "such url not exists: %s"
)

// CreateURL
// check cache for url request string
// if not exists, then get next unique short string by component
// save new url with unique short string in cache
// set response body with generated value
func (s URLService) CreateURL(ctx echo.Context, sourceURL string) (string, error) {
	var err error
	defer func(start time.Time) {
		if err != nil {
			log.Errorf(CreateURIErr, err.Error())
		}
		// TODO add RT metric here
		fmt.Println(time.Since(start))
	}(time.Now())

	key := sourceURL
	if URL, ok := s.cacheStore.Get(key); ok { // if user ask URL that already exists
		if v, ok2 := URL.(string); ok2 {
			return v, nil
		}

		return "", fmt.Errorf("cannot cast to string extracted value from db : %v", URL)
	}

	var generatedShortURL string // TODO need id generator component

	for { // do until getting unique string which not exists in cache (to exclude collisions)
		generatedShortURL = utils.GetUniqueString(GeneratedURLLen)
		if _, ok := s.shortURLStore.Get(generatedShortURL); !ok {
			break
		}
	}
	s.shortURLStore.Set(generatedShortURL, key)
	s.cacheStore.Set(key, generatedShortURL) // TODO how not store all records x2 times url <> short_url & short_url <> url ?

	return generatedShortURL, nil
}

func (s URLService) GetURL(ctx echo.Context, sourceURL string) (string, error) {
	key := sourceURL
	if URL, ok := s.shortURLStore.Get(key); ok { // if user ask URL that already exists
		if v, ok2 := URL.(string); ok2 {
			return v, nil
		}

		return "", fmt.Errorf("cannot cast to string extracted value from db : %v", URL)
	}

	return "", fmt.Errorf(SuchURLNotExists, key)
}
