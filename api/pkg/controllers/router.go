package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter returns new router
func NewRouter(h *Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", h.healthCheckHandler)
	e.GET("/github/auth/url", h.githubURL)
	e.GET("/github/auth/callback", h.githubCallback)

	return e
}

func (h *Handler) healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
