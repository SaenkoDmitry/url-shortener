package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"url-shortener/internal/api"
	"url-shortener/internal/cache"
	"url-shortener/internal/service"

	"github.com/labstack/echo-contrib/prometheus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiKeyCacheSize = 32 * 1024   // up to 32 k users
	URLCacheSize    = 1024 * 1024 // up to 1 million URL records // TODO limit requests by user api key
)

func main() {
	rand.Seed(time.Now().UnixNano())

	e := echo.New()

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	apiKeyCache, err := cache.NewCache(apiKeyCacheSize)
	mustInit(err)

	URLCache, err := cache.NewCache(URLCacheSize)
	mustInit(err)

	shortURLCache, err := cache.NewCache(URLCacheSize)
	mustInit(err)
	// ... some another stores or component

	userService := service.NewUserService(apiKeyCache)
	uriService := service.NewURLService(URLCache, shortURLCache)

	// Initialize api
	h := api.NewHandler(userService, uriService)

	e.POST("/api/user/api-key", h.CreateAPIKey)
	e.POST("/api/url", h.CreateURL)
	e.GET("/api/url", h.GetURL)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func mustInit(err error) {
	if err != nil {
		panic(err)
	}
}
