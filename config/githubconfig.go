package config

import (
	"context"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"log"
)

func CreateGithubClient() (*github.Client, context.Context) {
	err := ReadConfig("githubconfig")
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), ctx
}
