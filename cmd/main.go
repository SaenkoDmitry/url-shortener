package main

import (
	"log"
	"net/http"
	"url-shortener/internal/api"
	"url-shortener/internal/cache"
	"url-shortener/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	cacheSize = 128
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	uriCache, err := cache.NewCache(cacheSize)
	if err != nil {
		panic(err)
	}
	// ... some database

	// service 1
	userService := service.NewUserService(uriCache)

	// Initialize api
	h := api.NewHandler(userService)

	e.POST("/api/uri/create", h.CreateURI)
	e.POST("/api/user/api-key", h.CreateAPIKey)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
