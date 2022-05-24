package markdown

import (
	"bufio"
	"fmt"
	logger "github-stars/logging"
	"github-stars/schemas"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func WriteMarkDownFile(fileName string, allRepos []schemas.GitHubResponseField) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	markDownFile, err := os.Create(pwd + "/" + fileName)
	if err != nil {
		panic(err)
	}
	defer func(markDownFile *os.File) {
		err := markDownFile.Close()
		if err != nil {

		}
	}(markDownFile)

	writer := bufio.NewWriter(markDownFile)

	_, _ = fmt.Fprintln(writer, "# Starred Repositories"+"  ")
	_, _ = fmt.Fprintln(writer, "[How this generated?](../master/USAGE.md)"+"  ")
	_, _ = fmt.Fprintln(writer, "  ")
	_, _ = fmt.Fprintln(writer, "| Id 			| Name			| Description | Star Counts | Topics/Tags   | Last Updated 	|"+"  ")
	_, _ = fmt.Fprintln(writer, "| ----------- | ----------- 	| ----------- | ----------- | ----------- 	| -----------   |"+"  ")

	if err != nil {
		fmt.Println(err)
		err := markDownFile.Close()
		if err != nil {
			return
		}
		return
	}

	// Sort by Name
	//sort.Slice(allRepos[:], func(i, j int) bool {
	//	return allRepos[i].Name < allRepos[i].Name
	//})

	// Use of Reflection to sort all Repos
	By(Prop("Name", true)).Sort(allRepos)

	for index, getRepo := range allRepos {
		name := getRepo.Name
		fullName := getRepo.FullName
		lastUpdated := getRepo.LastUpdated

		description := getRepo.Description
		cloneUrl := getRepo.CloneUrl
		ownerName := getRepo.OwnerName
		starCount := getRepo.StarCount
		topics := strings.Join(getRepo.Topics, ", ")
		_, _ = fmt.Fprintln(writer, "|"+strconv.Itoa(index+1)+"|"+"["+name+"]"+"("+cloneUrl+")"+"|"+description+"|"+strconv.Itoa(starCount)+"|"+topics+"|"+lastUpdated+"|"+"  ")
		//_, err = writer.WriteString("|" + strconv.Itoa(index+1) + "|" + "[" + name + "]" + "(" + cloneUrl + ")" + "|" + description + "|" + strconv.Itoa(starCount) + "|" + topics + "|" + lastUpdated + "|" + "  " + "\n")

		//fmt.Printf("Id: %d, Name: %s, FullName: %s, Description: %s, CloneURL: %s, Owner: %s, StargazersCount: %d, LastUpdated: %s\n", index, name, fullName, description, cloneUrl, ownerName, starCount, lastUpdated)
		logger.WithFields(logrus.Fields{
			"Id":              index,
			"Name":            name,
			"FullName":        fullName,
			"Description":     description,
			"CloneURL":        cloneUrl,
			"Owner":           ownerName,
			"StargazersCount": starCount,
			"LastUpdated":     lastUpdated,
		}).Debug("")
	}
	_, _ = fmt.Fprintln(writer, "  ")
	err = writer.Flush()
	if err != nil {
		return
	}
}
