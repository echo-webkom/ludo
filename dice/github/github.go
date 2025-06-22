package github

import (
	"context"
	"fmt"

	"github.com/echo-webkom/ludo/dice/config"
	"github.com/google/go-github/v72/github"
)

type Branch struct {
	Name    string // Name of branch
	Owner   string // Owner of repo (eg. echo-webkom)
	Repo    string // Repo name
	Creator string // Username of user that created the branch
	PROpen  bool   // If a PR is open or not
	PRLink  string // URL to PR, else empty
	Merged  bool   // If the branch is merged
	Commits int    // Number of commits ahead of main
	Draft   bool   // If the PR is draft
}

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
func (c *Client) FetchBranchInfo(user, repoOwner, repoName, branchName string) (branch Branch, err error) {
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
		Head:  fmt.Sprintf("%s:%s", user, branchName),
	})
	if err != nil {
		return branch, err
	}

	prUrl, prCreator, prDraft, prOpen, prMerged := prInfo(prs)

	return Branch{
		Name:    branchName,
		Owner:   repoOwner,
		Repo:    repoName,
		Creator: prCreator,
		PRLink:  prUrl,
		PROpen:  prOpen,
		Draft:   prDraft,
		Merged:  prMerged,
		Commits: compare.GetAheadBy(),
	}, nil
}

func prInfo(prs []*github.PullRequest) (url, creator string, isDraft, isOpen, isMerged bool) {
	if len(prs) == 0 {
		return
	}
	pr := prs[0]
	return pr.GetURL(), pr.GetUser().GetName(), pr.GetDraft(), pr.ClosedAt != nil, pr.GetMerged()
}
