package github

import (
	"context"

	"github.com/google/go-github/v37/github"
)

func FetchCommits(username, repo string) ([]*github.RepositoryCommit, error) {
	client := NewClient()
	ctx := context.Background()

	var commits []*github.RepositoryCommit
	var err error

	if repo == "" {
		repos, _, err := client.Repositories.List(ctx, username, nil)
		if err != nil {
			return nil, err
		}

		for _, r := range repos {
			repoCommits, _, err := client.Repositories.ListCommits(ctx, username, r.GetName(), nil)
			if err != nil {
				return nil, err
			}
			commits = append(commits, repoCommits...)
		}
	} else {
		commits, _, err = client.Repositories.ListCommits(ctx, username, repo, nil)
		if err != nil {
			return nil, err
		}
	}

	return commits, nil
}
