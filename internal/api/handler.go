package api

import (
	"url-shortener/internal/service"
)

type Handler struct {
	us *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
		us: userService,
	}
}
