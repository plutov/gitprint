package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/plutov/gitprint/api/pkg/http/response"
)

type featuredRepo struct {
	Name        string `json:"name"`
	Size        string `json:"size"`
	Version     string `json:"version"`
	DownloadURL string `json:"download_url"`
}

var featuredRepos = []featuredRepo{
	{Name: "iancmcc/bingo", Size: "0.2MB", Version: "v0.2.0", DownloadURL: "https://gitprint.me/repos/bingo.pdf"},
	{Name: "adlio/schema", Size: "0.2MB", Version: "v1.4.0", DownloadURL: "https://gitprint.me/repos/schema.pdf"},
}

func (h *Handler) getPopularRepos(c echo.Context) error {
	return response.Ok(c, map[string]interface{}{
		"repos": featuredRepos,
	})
}
