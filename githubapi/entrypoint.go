package githubapi

import (
	"github-stars/markdown"
)

func CallGitHubAPIs(accessToken string, fileName string) {
	clientInfo := GetGitHubClient(accessToken)
	allRepos := clientInfo.GetGithubStarredRepoByUser()
	githubResponse := clientInfo.ParseGitHubApiResponse(allRepos)
	markdown.WriteMarkDownFile(fileName, githubResponse)
}
