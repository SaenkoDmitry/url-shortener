package service

import (
	"fmt"
	"url-shortener/internal/cache"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService struct {
	cacheStore cache.MyCache
}

func NewUserService(cacheStore cache.MyCache) *UserService {
	return &UserService{
		cacheStore: cacheStore,
	}
}

func (u *UserService) GetAPIKey(c echo.Context) (string, error) {
	remoteAddr := c.Request().RemoteAddr
	if APIKey, ok := u.cacheStore.Get(remoteAddr); ok {
		return APIKey.(string), nil
	}

	generatedKey, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("cannot generate uuid string for api key: %v", err)
	}

	if err1 := u.cacheStore.Set(remoteAddr, generatedKey); err1 != nil {
		return "", fmt.Errorf("cannot save api key: %v", err1.Error())
	}

	return generatedKey.String(), nil
}
