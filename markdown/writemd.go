package markdown

import (
	"bufio"
	"fmt"
	"github-stars/githubapi"
	"os"
	"strconv"
	"strings"
)

func WriteMarkDownFile(fileName string, allRepos []githubapi.GitHubResponseField) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//parentFolder := filepath.Dir(pwd)
	markDownFile, _ := os.Create(pwd + "/" + fileName)
	writer := bufio.NewWriter(markDownFile)

	_, _ = writer.WriteString("# Starred Repositories" + "  " + "\n\n")
	_, _ = writer.WriteString("[How this generated?](../master/USAGE.md)" + "  " + "\n\n")

	_, _ = writer.WriteString("| Id 			| Name			| Description | Star Counts | Topics/Tags   | Last Updated 	|" + "  " + "\n")
	_, _ = writer.WriteString("| ----------- | ----------- 	| ----------- | ----------- | ----------- 	| -----------   |" + "  " + "\n")

	if err != nil {
		fmt.Println(err)
		err := markDownFile.Close()
		if err != nil {
			return
		}
		return
	}

	for index, getRepo := range allRepos {
		name := getRepo.Name
		fullName := getRepo.FullName
		lastUpdated := getRepo.LastUpdated

		description := getRepo.Description
		cloneUrl := getRepo.CloneUrl
		ownerName := getRepo.OwnerName
		starCount := getRepo.StarCount
		topics := strings.Join(getRepo.Topics, ", ")
		_, err = writer.WriteString("|" + strconv.Itoa(index+1) + "|" + "[" + name + "]" + "(" + cloneUrl + ")" + "|" + description + "|" + strconv.Itoa(starCount) + "|" + topics + "|" + lastUpdated + "|" + "  " + "\n")

		fmt.Printf("Id: %d\tName: %s\tFullName: %s\tDescription: %s\tCloneURL: %s\tOwner: %s\tStargazersCount: %d\tLastUpdated: %s\n", index, name, fullName, description, cloneUrl, ownerName, starCount, lastUpdated)
	}
	writer.Flush()
}
