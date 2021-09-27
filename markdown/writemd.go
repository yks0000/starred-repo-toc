package markdown

import (
	"bufio"
	"fmt"
	"github.com/google/go-github/github"
	"os"
	"strconv"
	"strings"
)

func WriteMarkDownFile(fileName string, allRepos []*github.StarredRepository) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//parentFolder := filepath.Dir(pwd)
	markDownFile, _ := os.Create(pwd + "/" + fileName)
	writer := bufio.NewWriter(markDownFile)

	_, _ = writer.WriteString("# List of Starred Repository" + "  " + "\n\n")

	_, _ = writer.WriteString("| Id 			| Name			| Description | Star Counts | Last Updated 	|" + "  " + "\n")
	_, _ = writer.WriteString("| ----------- | ----------- 	| ----------- | ----------- | ----------- 	|" + "  " + "\n")

	if err != nil {
		fmt.Println(err)
		err := markDownFile.Close()
		if err != nil {
			return
		}
		return
	}

	for index, getRepo := range allRepos {
		repoDetails := getRepo.GetRepository()
		name := *repoDetails.Name
		fullName := *repoDetails.FullName

		year, month, day := repoDetails.PushedAt.Date()
		lastUpdated := strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)

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
		_, err = writer.WriteString("|" + strconv.Itoa(index+1) + "|" + "[" + name + "]" + "(" + cloneUrl + ")" + "|" + description + "|" + strconv.Itoa(starCount) + "|" + lastUpdated + "|" + "  " + "\n")

		fmt.Printf("Id: %d\tName: %s\tFullName: %s\tDescription: %s\tCloneURL: %s\tOwner: %s\tStargazersCount: %d\tLastUpdated: %s\n", index, name, fullName, description, cloneUrl, ownerName, starCount, lastUpdated)
	}
	writer.Flush()
}
