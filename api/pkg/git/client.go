package git

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/v65/github"
	"github.com/plutov/gitprint/api/pkg/log"
)

type Client struct {
	client   *github.Client
	reposDir string
}

type Org struct {
	Name string
}

type Repo struct {
	Name string
}

type GetContentsResult struct {
	Files int64
	Dirs  int64
}

func NewClient(accessToken string) *Client {
	return &Client{
		client:   github.NewClient(nil).WithAuthToken(accessToken),
		reposDir: os.Getenv("GITHUB_REPOS_DIR"),
	}
}

func (c *Client) GetUserOrgs() ([]Org, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("getting user orgs")

	ghOrgs, _, err := c.client.Organizations.List(ctx, "", nil)
	if err != nil {
		log.WithError(err).Error("failed to get user orgs")
		return nil, err
	}

	orgs := make([]Org, len(ghOrgs))
	for i, ghOrg := range ghOrgs {
		orgs[i] = Org{
			Name: ghOrg.GetLogin(),
		}
	}

	return orgs, nil
}

func (c *Client) GetUserRepos() ([]Repo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("getting user repos")

	ghRepos, _, err := c.client.Repositories.ListByAuthenticatedUser(ctx, nil)
	if err != nil {
		log.WithError(err).Error("failed to get user orgs")
		return nil, err
	}

	repos := make([]Repo, len(ghRepos))
	for i, ghRepo := range ghRepos {
		repos[i] = Repo{
			Name: ghRepo.GetName(),
		}
	}

	return repos, nil
}

func (c *Client) GetOrgRepos(org string) ([]Repo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("getting user repos")

	ghRepos, _, err := c.client.Repositories.ListByOrg(ctx, org, nil)
	if err != nil {
		log.WithError(err).Error("failed to get user orgs")
		return nil, err
	}

	repos := make([]Repo, len(ghRepos))
	for i, ghRepo := range ghRepos {
		repos[i] = Repo{
			Name: ghRepo.GetName(),
		}
	}

	return repos, nil
}

func (c *Client) DownloadRepo(owner string, repo string, ref string) (*GetContentsResult, error) {
	logCtx := log.With("owner", owner, "repo", repo, "ref", ref)
	logCtx.Info("downloading repo")

	res, err := c.getContents(owner, repo, ref, "")
	if err != nil {
		logCtx.WithError(err).Error("failed to download repo")
		return nil, err
	}

	logCtx.Info("repo downloaded")
	return res, nil
}

func (c *Client) getContents(owner string, repo string, ref string, path string) (*GetContentsResult, error) {
	res := &GetContentsResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, directoryContent, _, err := c.client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{
		Ref: ref,
	})
	if err != nil {
		return nil, err
	}

	for _, item := range directoryContent {
		switch *item.Type {
		case "file":
			localPath := filepath.Join(c.reposDir, *item.Path)
			if err := c.downloadContents(owner, repo, ref, item, localPath); err != nil {
				return nil, fmt.Errorf("failed to download file %s: %w", *item.Path, err)
			}
			res.Files++
		case "dir":
			dirRes, err := c.getContents(owner, repo, ref, *item.Path)
			if err != nil {
				return nil, fmt.Errorf("failed to download dir %s: %w", *item.Path, err)
			}

			res.Dirs++
			res.Files += dirRes.Files
			res.Dirs += dirRes.Dirs
		}
	}

	return res, nil
}

func (c *Client) downloadContents(owner string, repo string, ref string, content *github.RepositoryContent, localPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	rc, _, err := c.client.Repositories.DownloadContents(ctx, owner, repo, *content.Path, &github.RepositoryContentGetOptions{
		Ref: ref,
	})
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer rc.Close()

	b, err := io.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(localPath), 0755)
	if err != nil {
		return fmt.Errorf("failed to create a local dir %s: %w", filepath.Dir(localPath), err)
	}

	f, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create a file: %w", err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("failed to write a file: %w", err)
	}

	return nil
}
