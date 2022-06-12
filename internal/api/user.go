package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateAPIKey(ctx echo.Context) error {
	apiKey, err := h.us.GetAPIKey(ctx)
	if err != nil {
		// add logger
		return err
	}

	return ctx.JSON(http.StatusOK, apiKey)
}
