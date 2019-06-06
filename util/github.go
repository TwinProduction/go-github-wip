package util

import (
	"context"
	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
	"os"
)

var (
	githubClient *github.Client
)

func GetGithubClient() (*github.Client, context.Context) {
	if githubClient == nil {
		githubClient = createGithubClient()
	}
	return githubClient, context.Background()
}

func createGithubClient() *github.Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	client := oauth2.NewClient(ctx, tokenSource)
	return github.NewClient(client)
}
