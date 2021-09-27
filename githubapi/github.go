package githubapi

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
)

var allRepos []*github.StarredRepository

func GetGitHubClient(accessToken string) (*github.Client, context.Context) {
	ctx := context.Background()
	secureTokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, secureTokenSource)
	client := github.NewClient(tc)
	return client, ctx
}

func GetGithubStarredRepoByUser(client *github.Client, context context.Context) []*github.StarredRepository {
	user, _, err := client.Users.Get(context, "")
	if err != nil {
		log.Panicln("Error while making authenticated call to github", err)
	}

	activityListStarredOptions := &github.ActivityListStarredOptions{ListOptions: github.ListOptions{PerPage: 100}}

	for {
		repos, resp, err := client.Activity.ListStarred(context, *user.Login, activityListStarredOptions)
		if err != nil {
			log.Panicln("Error while making authenticated call to github", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		} else {
			log.Printf("Loading another page. Page loaded: %d", resp.NextPage)
		}
		activityListStarredOptions.Page = resp.NextPage
	}
	return allRepos

}
