package git

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/v65/github"
	"github.com/plutov/gitprint/api/pkg/log"
)

type Client struct {
	accessToken string
	client      *github.Client
	reposDir    string
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		client:      github.NewClient(nil).WithAuthToken(accessToken),
		reposDir:    os.Getenv("GITHUB_REPOS_DIR"),
	}
}

func (c *Client) GetCurrentUser() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("getting current user")

	me, _, err := c.client.Users.Get(ctx, "")
	if err != nil {
		log.WithError(err).Error("failed to get current user")
		return nil, err
	}

	return &User{
		Username:    me.GetLogin(),
		AccessToken: c.accessToken,
	}, nil
}

func (c *Client) GetTopContributors(owner string, repo string) ([]*github.Contributor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("getting repo contributors")

	opts := &github.ListContributorsOptions{
		Anon: "false",
	}
	contributors, _, err := c.client.Repositories.ListContributors(ctx, owner, repo, opts)
	if err != nil {
		log.WithError(err).Error("failed to get repo contributors")
		return nil, err
	}

	if len(contributors) > 50 {
		contributors = contributors[:50]
	}

	return contributors, nil
}

func (c *Client) GetRepo(owner string, repo string) (*github.Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("getting repo")

	repository, _, err := c.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		log.WithError(err).Error("failed to get repo")
		return nil, err
	}

	return repository, nil
}

type DownloadRepoResult struct {
	Ref        string
	OutputFile string
}

func (c *Client) DownloadRepo(owner string, repo string, ref string) (*DownloadRepoResult, error) {
	logCtx := log.With("owner", owner, "repo", repo)

	// Get latest commit if ref is empty
	if ref == "" {
		var shaErr error
		ref, shaErr = c.GetLatestCommitSHA(owner, repo)
		if shaErr != nil {
			logCtx.WithError(shaErr).Error("failed to get latest commit")
			return nil, shaErr
		}
	}

	logCtx = logCtx.With("ref", ref)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logCtx.Info("getting archive url")
	opts := &github.RepositoryContentGetOptions{}
	opts.Ref = ref
	url, _, err := c.client.Repositories.GetArchiveLink(ctx, owner, repo, github.Tarball, opts, 1)
	if err != nil {
		logCtx.WithError(err).Error("failed to get archive link")
		return nil, err
	}
	logCtx.With("url", url.String())

	outputFile := filepath.Join(c.reposDir, owner, repo, ref+".tar.gz")
	// Create output directory
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		logCtx.WithError(err).Error("failed to create output directory")
		return nil, err
	}

	out, err := os.Create(outputFile)
	if err != nil {
		logCtx.WithError(err).Error("failed to create output file")
		return nil, err
	}
	defer out.Close()

	logCtx.Info("downloading archive")
	client := http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Get(url.String())
	if err != nil {
		logCtx.WithError(err).Error("failed to download repo")
		return nil, err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		logCtx.Errorf("non-200 response code: %d", resp.StatusCode)
		return nil, fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logCtx.WithError(err).Error("failed to write to file")
		return nil, err
	}

	logCtx.Info("repo downloaded")
	return &DownloadRepoResult{
		Ref:        ref,
		OutputFile: outputFile,
	}, nil
}

func (c *Client) GetLatestCommitSHA(owner string, repo string) (string, error) {
	logCtx := log.With("owner", owner, "repo", repo)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logCtx.Info("getting latest commit")

	opts := &github.CommitsListOptions{}
	opts.PerPage = 1
	commits, _, err := c.client.Repositories.ListCommits(ctx, owner, repo, opts)
	if err != nil {
		logCtx.WithError(err).Error("failed to get latest commit")
		return "", err
	}
	if len(commits) == 0 {
		logCtx.Error("no commits found")
		return "", fmt.Errorf("no commits found")
	}

	return *commits[0].SHA, nil
}
