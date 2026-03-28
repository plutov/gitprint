package controllers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/plutov/gitprint/api/pkg/auth"
	"github.com/plutov/gitprint/api/pkg/log"
)

// NewRouter returns new router
func NewRouter(h *Handler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// open endpoints
	e.GET("/", h.healthCheckHandler)
	e.GET("/github/auth/url", h.githubURL)
	e.GET("/github/auth/callback", h.githubCallback)
	e.GET("/files", h.downloadExportFile)
	e.GET("/repos/popular", h.getPopularRepos)

	// restricted endpoints
	private := e.Group("/private")
	private.Use(h.authMiddleware)
	private.GET("/github/repo/download", h.downloadRepo)
	private.GET("/github/repo/generate", h.generate)

	return e
}

func (h *Handler) healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

func (h *Handler) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtToken := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
		if jwtToken == "" {
			return c.JSON(http.StatusUnauthorized, "missing jwt token")
		}

		session, err := auth.ReadJWTClaims(jwtToken)
		if err != nil {
			log.WithError(err).Error("failed to read jwt claims")
			return c.JSON(http.StatusUnauthorized, "invalid jwt token")
		}

		c.Set("user", session.User)

		return next(c)
	}
}
