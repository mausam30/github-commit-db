package github

import (
	"context"
	"os"
	"strings"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

type CommitWithRepo struct {
	*github.RepositoryCommit
	Repository *github.Repository `json:"repository"`
}

func FetchCommits(username, repo string) ([]CommitWithRepo, error) {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	var allCommits []CommitWithRepo
	if repo != "" {
		// Fetch commits for the specified repository
		commits, _, err := client.Repositories.ListCommits(ctx, username, repo, nil)
		if err != nil {
			return nil, err
		}
		for _, commit := range commits {
			allCommits = append(allCommits, CommitWithRepo{commit, &github.Repository{Name: github.String(repo)}})
		}
	} else {
		// Fetch repositories for the user
		repos, _, err := client.Repositories.List(ctx, username, nil)
		if err != nil {
			return nil, err
		}
		for _, repository := range repos {
			commits, _, err := client.Repositories.ListCommits(ctx, username, repository.GetName(), nil)
			if err != nil {
				if strings.Contains(err.Error(), "409") { // Handle empty repositories
					continue
				}
				return nil, err
			}
			for _, commit := range commits {
				allCommits = append(allCommits, CommitWithRepo{commit, repository})
			}
		}
	}

	return allCommits, nil
}
