package service

import (
	"fmt"
	"net/http"
	"time"
	"url-shortener/internal/cache"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	expires = 24 * time.Hour
)

type UserService struct {
	cacheStore cache.MyCache
}

func NewUserService(cacheStore cache.MyCache) *UserService {
	return &UserService{
		cacheStore: cacheStore,
	}
}

func (u *UserService) GetAPIKey(ctx echo.Context) (string, error) {
	remoteAddr := ctx.Request().RemoteAddr
	key := remoteAddr

	if APIKey, ok := u.cacheStore.Get(key); ok {
		return APIKey.(string), nil
	}

	uniqueUUID, err := uuid.NewUUID() // need special module for generating unique api keys
	if err != nil {
		return "", fmt.Errorf("cannot generate uuid string for api key: %v", err)
	}

	userAPIKey := uniqueUUID.String()

	u.cacheStore.Set(key, userAPIKey)
	setAPIKeyCookie(ctx, userAPIKey)

	return userAPIKey, nil
}

func setAPIKeyCookie(ctx echo.Context, key string) {
	cookie := new(http.Cookie)
	cookie.Name = "api-key"
	cookie.Value = key
	cookie.Expires = time.Now().Add(expires)
	ctx.SetCookie(cookie)
}
