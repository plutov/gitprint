package controllers

import (
	"github.com/google/go-github/v65/github"
	"github.com/labstack/echo/v4"
	"github.com/plutov/gitprint/api/pkg/builder"
	"github.com/plutov/gitprint/api/pkg/files"
	"github.com/plutov/gitprint/api/pkg/git"
	"github.com/plutov/gitprint/api/pkg/http/response"
)

func (h *Handler) getOrgs(c echo.Context) error {
	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	me, err := ghClient.GetCurrentUser()
	if err != nil {
		return response.InternalError(c, "unable to get current user")
	}

	orgs, err := ghClient.GetUserOrgs()
	if err != nil {
		return response.InternalError(c, "unable to get user orgs")
	}

	combined := []string{me.Username}
	for _, o := range orgs {
		combined = append(combined, o.GetLogin())
	}

	return response.Ok(c, echo.Map{
		"orgs": combined,
	})
}

func (h *Handler) getRepos(c echo.Context) error {
	owner := c.QueryParam("owner")

	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	var (
		repos []*github.Repository
		err   error
	)
	if owner == "" || owner == user.Username {
		repos, err = ghClient.GetOrgRepos(owner)
	} else {
		repos, err = ghClient.GetOrgRepos(owner)
	}
	if err != nil {
		return response.InternalError(c, "unable to get repos")
	}

	return response.Ok(c, echo.Map{
		"repos": repos,
	})
}

func (h *Handler) downloadRepo(c echo.Context) error {
	owner := c.QueryParam("owner")
	repo := c.QueryParam("repo")
	ref := c.QueryParam("ref")

	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	res, err := ghClient.DownloadRepo(owner, repo, ref)
	if err != nil {
		return response.InternalError(c, "unable to download repo")
	}

	extracted, err := files.ExtractAndFilterFiles(res.OutputFile)
	if err != nil {
		return response.InternalError(c, "unable to extract and filter files")
	}

	return response.Ok(c, extracted)
}

func (h *Handler) generate(c echo.Context) error {
	owner := c.QueryParam("owner")
	repo := c.QueryParam("repo")
	outputDir := c.QueryParam("output_dir")

	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	repository, err := ghClient.GetRepo(owner, repo)
	if err != nil {
		return response.InternalError(c, "unable to get repo")
	}

	contributors, err := ghClient.GetTopContributors(owner, repo)
	if err != nil {
		return response.InternalError(c, "unable to get repo contributors")
	}

	commitSHA, err := ghClient.GetLatestCommitSHA(owner, repo)
	if err != nil {
		return response.InternalError(c, "unable to get last commit SHA")
	}

	doc, err := builder.GenerateDocument(repository, contributors, commitSHA, outputDir)
	if err != nil {
		return response.InternalError(c, "unable to generate document")
	}

	return response.Ok(c, doc)
}
