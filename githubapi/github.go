package githubapi

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"strconv"
	"strings"
)

var allRepos []*github.StarredRepository

type GitHubResponseField struct {
	Name string
	FullName string
	Description string
	CloneUrl string
	OwnerName string
	StarCount int
	LastUpdated string
	Topics []string
}

var githubResponseField []GitHubResponseField

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


func ParseGitHubApiResponse(allRepos []*github.StarredRepository, client *github.Client, context context.Context) []GitHubResponseField {
	for _, getRepo := range allRepos {
		repoDetails := getRepo.GetRepository()
		name := *repoDetails.Name
		fullName := *repoDetails.FullName
		defaultBranch := *repoDetails.DefaultBranch

		description := "-"
		if repoDetails.Description != nil {
			description = strings.Replace(*repoDetails.Description, "|", " ", -1)
		}
		cloneUrl := *repoDetails.CloneURL
		ownerName := "-"
		if repoDetails.Owner != nil && repoDetails.Owner.Login != nil {
			ownerName = *repoDetails.Owner.Login
		}
		starCount := *repoDetails.StargazersCount

		lastUpdatedChannel := make(chan string)
		go GetDefaultBranchDetails(client, context, name, ownerName, defaultBranch, lastUpdatedChannel)
		lastUpdated := <- lastUpdatedChannel

		channels := make(chan []string)
		go GetGitHubRepoTopics(client, context, name, ownerName, channels)
		topics := <- channels
		githubResponseField = append(githubResponseField, GitHubResponseField{
			Name: name,
			FullName: fullName,
			Description: description,
			CloneUrl: cloneUrl,
			OwnerName: ownerName,
			StarCount: starCount,
			LastUpdated: lastUpdated,
			Topics: topics,
		})
	}
	return githubResponseField
}

func GetDefaultBranchDetails(client *github.Client, context context.Context, repoName string, ownerName string, branchName string, channel chan string) {
	branch, _, err := client.Repositories.GetBranch(context, ownerName, repoName, branchName)
	if err != nil {
		log.Println(err)
	}
	year, month, day := branch.GetCommit().Commit.Committer.GetDate().Date()
	channel <- strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)

}

func GetGitHubRepoTopics(client *github.Client, context context.Context, repoName string, ownerName string, channel chan []string) {
	log.Println("Getting topics tag for repo", repoName)
	topics, _, err := client.Repositories.ListAllTopics(context, ownerName, repoName)
	if err != nil {
		log.Println(err)
	}
	channel <- topics
}