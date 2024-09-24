package git

import (
	"context"
	"time"

	"github.com/google/go-github/v65/github"
	"github.com/plutov/gitprint/api/pkg/log"
)

type Client struct {
	client *github.Client
}

type Org struct {
	Name string
}

type Repo struct {
	Name string
}

func NewClient(accessToken string) *Client {
	return &Client{
		client: github.NewClient(nil).WithAuthToken(accessToken),
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
