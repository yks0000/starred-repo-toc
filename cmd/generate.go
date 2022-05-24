package cmd

import (
	"fmt"
	"github-stars/githubapi"
	"github.com/spf13/cobra"
)

var fileName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate MD with Star Information.",
	Run: func(cmd *cobra.Command, args []string) {
		githubapi.CallGitHubAPIs(accessToken, fileName)
		fmt.Println("Done!!!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "README.md", "MarkDown File Name")
}
