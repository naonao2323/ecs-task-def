package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	githubClient *github.Client
	owner        string
	repo         string
}

type Github interface {
	CreatePullRequest(ctx context.Context, headBranch string) error
}

func NewGithubClient(ctx context.Context, token string, owner string, repo string) Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	return &GithubClient{githubClient: githubClient, owner: owner, repo: repo}
}

func (g *GithubClient) CreatePullRequest(ctx context.Context, headBranch string) error {
	pr := &github.NewPullRequest{
		Title: github.String("PR Title"),
		Head:  github.String(headBranch),
		Base:  github.String("main"),
		Body:  github.String("Tetst PR description."),
	}
	_, _, err := g.githubClient.PullRequests.Create(ctx, g.owner, g.repo, pr)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
