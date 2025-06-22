package github

import (
	"context"
	"fmt"

	"github.com/echo-webkom/ludo/dice/config"
	"github.com/google/go-github/v72/github"
)

type Branch struct {
	Name    string       // Name of branch
	Owner   string       // Owner of repo (eg. echo-webkom)
	Repo    string       // Repo name
	Commits int          // Number of commits ahead of main
	PR      *PullRequest // nil if none are open
}

type PullRequest struct {
	Creator string // Username of user that created the pr
	Open    bool   // If the pr is open or not
	Url     string // URL to pr
	Draft   bool   // If the pr is draft
	Merged  bool   // If the pr is merged
}

// Get html url to branch at github
func (b Branch) URL() string {
	return fmt.Sprintf("https://github.com/%s/%s/tree/%s", b.Repo, b.Owner, b.Name)
}

type Client struct {
	client *github.Client
}

func New(config *config.Config) *Client {
	client := github.NewClient(nil).WithAuthToken(config.GitHubAuthToken)
	return &Client{client}
}

// Fetch info on given branch in repo with owner. Returns error if branch does
// not exist or the request failed.
func (c *Client) FetchBranchInfo(repoOwner, repoName, branchName string) (branch Branch, err error) {
	ctx := context.Background()
	repo, _, err := c.client.Repositories.Get(ctx, repoOwner, repoName)
	if err != nil {
		return branch, err
	}

	compare, _, err := c.client.Repositories.CompareCommits(ctx, repoOwner, repoName, repo.GetDefaultBranch(), branchName, nil)
	if err != nil {
		return branch, err
	}

	prs, _, err := c.client.PullRequests.List(ctx, repoOwner, repoName, &github.PullRequestListOptions{
		State: "open",
		Head:  fmt.Sprintf("%s:%s", repoOwner, branchName),
	})
	if err != nil {
		return branch, err
	}

	return Branch{
		Name:    branchName,
		Owner:   repoOwner,
		Repo:    repoName,
		Commits: compare.GetAheadBy(),
		PR:      prInfo(prs),
	}, nil
}

func prInfo(prs []*github.PullRequest) *PullRequest {
	if len(prs) == 0 {
		return nil
	}

	pr := prs[0]
	return &PullRequest{
		Url:     pr.GetHTMLURL(),
		Creator: pr.GetUser().GetLogin(),
		Draft:   pr.GetDraft(),
		Open:    true,
		Merged:  pr.GetMerged(),
	}
}
