package api

import (
	"fmt"
	"net/http"
	"net/url"
	"url-shortener/internal/models"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateURL(ctx echo.Context) error {
	req, err := validateRequest(ctx)
	if err != nil {
		return err
	}

	shortenedURL, err := h.url.CreateURL(ctx, req.URL)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &models.URLResponse{
		URL: shortenedURL,
	})
}

func (h *Handler) GetURL(ctx echo.Context) error {
	req := &models.URLRequest{}
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	shortenedURL, err := h.url.GetURL(ctx, req.URL)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, &models.URLResponse{
		URL: shortenedURL,
	})
}

func validateRequest(ctx echo.Context) (*models.URLRequest, error) {
	request := &models.URLRequest{}
	if err := ctx.Bind(request); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// extract validation on another step if needed
	_, err := url.ParseRequestURI(request.URL)
	if err != nil {
		err = fmt.Errorf("invalid input url: %s", err.Error())
		return nil, ctx.String(http.StatusBadRequest, err.Error())
	}

	return request, nil
}
