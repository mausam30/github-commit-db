package github

import (
	"context"
	"os"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

func NewClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
