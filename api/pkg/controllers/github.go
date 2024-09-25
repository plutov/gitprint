package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/plutov/gitprint/api/pkg/git"
	"github.com/plutov/gitprint/api/pkg/http/response"
)

func (h *Handler) githubURL(c echo.Context) error {
	url := h.Services.GithubAuth.GetRedirectURL()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handler) githubCallback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	token, err := h.Services.GithubAuth.GetAccessToken(code, state)
	if err != nil {
		return response.BadRequestDefaultMessage(c)
	}

	// TODO: move to another endpoint
	ghClient := git.NewClient(token)
	orgs, _ := ghClient.GetUserOrgs()
	fmt.Println(orgs)

	repos, _ := ghClient.GetUserRepos()
	fmt.Println(repos)

	repos, _ = ghClient.GetOrgRepos("12traits")
	fmt.Println(repos)

	res, _ := ghClient.DownloadRepo("plutov", "formulosity", "main")
	fmt.Println(res)

	return response.Ok(c, token)
}
