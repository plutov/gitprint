package controllers

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/plutov/gitprint/api/pkg/builder"
	"github.com/plutov/gitprint/api/pkg/files"
	"github.com/plutov/gitprint/api/pkg/git"
	"github.com/plutov/gitprint/api/pkg/http/response"
)

type repoRequest struct {
	Repo string `query:"repo"`
	Ref  string `query:"ref"`
}

func (r *repoRequest) Validate() error {
	parts := strings.Split(r.Repo, "/")
	if len(parts) != 2 {
		return errors.New("invalid repo")
	}

	if parts[0] == "" || parts[1] == "" {
		return errors.New("invalid repo")
	}

	return nil
}

func (r *repoRequest) GetOwnerAndRepo() (string, string) {
	parts := strings.Split(r.Repo, "/")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func (h *Handler) downloadRepo(c echo.Context) error {
	req := new(repoRequest)
	c.Bind(req)
	if err := req.Validate(); err != nil {
		return response.BadRequest(c, err.Error())
	}
	owner, repo := req.GetOwnerAndRepo()

	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	res, err := ghClient.DownloadRepo(owner, repo, req.Ref)
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
	req := new(repoRequest)
	c.Bind(req)
	if err := req.Validate(); err != nil {
		return response.BadRequest(c, err.Error())
	}
	owner, repo := req.GetOwnerAndRepo()
	exportID := c.QueryParam("export_id")
	if err := files.ValidateExportID(exportID); err != nil {
		return response.BadRequest(c, err.Error())
	}

	user := c.Get("user").(*git.User)
	ghClient := git.NewClient(user.AccessToken)

	repository, err := ghClient.GetRepo(owner, repo)
	if err != nil {
		return response.InternalError(c, "unable to find repo")
	}

	contributors, err := ghClient.GetTopContributors(owner, repo)
	if err != nil {
		return response.InternalError(c, "unable to get repo contributors")
	}

	if req.Ref == "" {
		var shaErr error
		req.Ref, shaErr = ghClient.GetLatestCommitSHA(owner, repo)
		if shaErr != nil {
			return response.InternalError(c, "unable to get last commit sha")
		}
	}

	doc, err := builder.GenerateDocument(repository, contributors, req.Ref, exportID)
	if err != nil {
		return response.InternalError(c, "unable to generate a document")
	}

	htmlOut, err := builder.GenerateAndSaveHTMLFile(doc, exportID)
	if err != nil {
		return response.InternalError(c, "unable to save html file")
	}

	_, err = builder.GenerateAndSavePDFFile(htmlOut, exportID)
	if err != nil {
		return response.InternalError(c, "unable to save pdf file")
	}

	return response.Ok(c, "ok")
}
