package api

import (
	"url-shortener/internal/service"
)

type Handler struct {
	us  *service.UserService
	url *service.URLService
}

func NewHandler(userService *service.UserService, uriService *service.URLService) *Handler {
	return &Handler{
		us:  userService,
		url: uriService,
	}
}
