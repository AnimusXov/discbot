package bot

import (
	"context"
	"github.com/google/go-github/v47/github"
	"strconv"
	"strings"
)

func listAllForks(owner string, repo string, client *github.Client, ctx context.Context) []*github.Repository {

	opt := &github.RepositoryListForksOptions{
		Sort: "newest",
	}
	forks, response, err := client.Repositories.ListForks(ctx, owner, repo, opt)
	if err != nil {
		return nil
	}
	if response.StatusCode != 200 {
		return nil
	}

	return forks
}

func compareBranches(client *github.Client, ctx context.Context, repo *github.Repository, owner string) (string, bool) {
	commits, response, err := client.Repositories.CompareCommits(
		ctx,
		repo.GetOwner().GetLogin(),
		repo.GetName(),
		owner+":"+"master",
		repo.GetOwner().GetLogin()+":"+"master",
		&github.ListOptions{
			Page:    1,
			PerPage: 1,
		})
	if err != nil {
		return err.Error(), false
	}
	if response.StatusCode != 200 {
		return strconv.Itoa(response.StatusCode), true
	}
	if strings.Compare(commits.GetStatus(), "ahead") == 0 {
		return "Ahead by: " + strconv.Itoa(commits.GetAheadBy()), true
	}
	if strings.Compare(commits.GetStatus(), "behind") == 0 {
		return "Behind by: " + strconv.Itoa(commits.GetBehindBy()), false
	}
	return "no difference", false
}
