package api

import "github.com/labstack/echo/v4"

func (h *Handler) CreateAPIKey(c echo.Context) error {
	apiKey, err := h.us.GetAPIKey(c)
	_ = apiKey
	return err
}
