package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	githubClient *github.Client
	logger       *zap.Logger
	owner        string
	repo         string
}

type Github interface {
	CreatePullRequest(ctx context.Context, headBranch string, tag string) error
}

func NewGithubClient(ctx context.Context, logger *zap.Logger, token string, owner string, repo string) Github {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	return &GithubClient{githubClient: githubClient, logger: logger, owner: owner, repo: repo}
}

func (g *GithubClient) CreatePullRequest(ctx context.Context, headBranch string, tag string) error {
	pr := &github.NewPullRequest{
		Title: github.String(fmt.Sprintf("ecs tag %s", tag)),
		Head:  github.String(headBranch),
		Base:  github.String("main"),
		Body:  github.String(""),
	}
	_, _, err := g.githubClient.PullRequests.Create(ctx, g.owner, g.repo, pr)
	if err != nil {
		g.logger.Error("fail to create pull request", zap.Error(err))
		return err
	}
	return nil
}
