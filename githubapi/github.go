package githubapi

import (
	"context"
	logger "github-stars/logging"
	"github-stars/schemas"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strconv"
	"strings"
	"sync"
	"time"
)

type GithubClientInformation struct {
	client  *github.Client
	context context.Context
}

var allRepos []*github.StarredRepository
var githubResponseField []schemas.GitHubResponseField
var pageLoadSleepTime = 2 * time.Second

func GetGitHubClient(accessToken string) GithubClientInformation {
	ctx := context.Background()
	secureTokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, secureTokenSource)
	client := github.NewClient(tc)
	return GithubClientInformation{
		client:  client,
		context: ctx,
	}
}

func (clientInfo GithubClientInformation) GetGithubStarredRepoByUser() []*github.StarredRepository {
	user, _, err := clientInfo.client.Users.Get(clientInfo.context, "")
	if err != nil {
		logger.Panic("Error while making authenticated call to github: ", err.Error())
	}

	activityListStarredOptions := &github.ActivityListStarredOptions{ListOptions: github.ListOptions{PerPage: 100}}

	for {
		repos, resp, err := clientInfo.client.Activity.ListStarred(clientInfo.context, *user.Login, activityListStarredOptions)
		if err != nil {
			logger.Panic("Error while making authenticated call to github: ", err.Error())
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		} else {
			logger.Info("Loading another page in 2 seconds. Page loaded: ", resp.NextPage)
		}
		activityListStarredOptions.Page = resp.NextPage
	}
	return allRepos

}

func (clientInfo *GithubClientInformation) ParseGitHubApiResponse(allRepos []*github.StarredRepository) []schemas.GitHubResponseField {
	wg := sync.WaitGroup{}
	wg.Add(len(allRepos))
	getData := func(getRepo *github.StarredRepository, wg *sync.WaitGroup) []schemas.GitHubResponseField {
		defer wg.Done()
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

		_, r, err := clientInfo.client.Repositories.GetByID(clientInfo.context, *repoDetails.ID)
		if r.StatusCode == 404 {
			clientInfo.UnStarDeleteGitHubRepo(name, ownerName)
			return nil
		}
		if err != nil {
			return nil
		}

		time.Sleep(pageLoadSleepTime)
		lastUpdated := clientInfo.GetDefaultBranchDetails(name, ownerName, defaultBranch)

		channels := make(chan []string)
		go clientInfo.GetGitHubRepoTopics(name, ownerName, channels)
		topics := <-channels
		githubResponseField = append(githubResponseField, schemas.GitHubResponseField{
			Name:        name,
			FullName:    fullName,
			Description: description,
			CloneUrl:    cloneUrl,
			OwnerName:   ownerName,
			StarCount:   starCount,
			LastUpdated: lastUpdated,
			Topics:      topics,
		})

		return githubResponseField
	}

	for _, getRepo := range allRepos {

		go getData(getRepo, &wg)

	}

	wg.Wait()
	return githubResponseField
}

func (clientInfo *GithubClientInformation) GetDefaultBranchDetails(repoName string, ownerName string, branchName string) string {
	branch, _, err := clientInfo.client.Repositories.GetBranch(clientInfo.context, ownerName, repoName, branchName)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	year, month, day := branch.GetCommit().Commit.Committer.GetDate().Date()
	return strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)

}

func (clientInfo *GithubClientInformation) GetGitHubRepoTopics(repoName string, ownerName string, channel chan []string) {
	time.Sleep(pageLoadSleepTime)
	logger.Info("Getting topics tag for repo: ", repoName)
	topics, _, err := clientInfo.client.Repositories.ListAllTopics(clientInfo.context, ownerName, repoName)
	if err != nil {
		logger.Error(err.Error())
	}
	channel <- topics
}

func (clientInfo *GithubClientInformation) UnStarDeleteGitHubRepo(repoName string, ownerName string) {
	unstar, err := clientInfo.client.Activity.Unstar(clientInfo.context, ownerName, repoName)
	logger.Info("Un starred deleted github repository ", repoName, unstar.StatusCode)
	if err != nil {
		return
	}
}
