package cmd

import (
	"fmt"
	"log"

	"github.com/born2ngopi/orca/generator"
	"github.com/born2ngopi/orca/git"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of orca",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Orca v0.1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "orca",
	Short: "Orca is a tool to help you write a commit message based on the changes in your project",
	Long: `
Orca is a tool to help you write a commit message based on the changes in your project,

This use local llm to generate commit message`,
	Run: func(cmd *cobra.Command, args []string) {

		model, _ := cmd.Flags().GetString("model")
		preview, _ := cmd.Flags().GetBool("preview")

		diffFiles, err := git.GetDiffFiles()
		if err != nil {
			log.Fatal(err)
		}

		var prompt = `
		Title: Generate commit message based on changes in a code block

Description: You are tasked with generating a commit message based on the changes made within a specific code block delimited by <commit></commit> tags in a Git diff.

Instructions:

Context: Imagine you are working with a Git repository containing code changes. Within the repository, there exists a file or files with code blocks enclosed by <commit></commit> tags.

Task: Your task is to examine the changes made within these <commit></commit> blocks and generate a concise and informative commit message describing these modifications.

Input:

Within the diff, you will identify the code blocks delimited by <commit></commit> tags.


`
		// var template = generator.Template{}
		for path, diff := range diffFiles {
			prompt += fmt.Sprintf("file: %s\ndif:\n%s\n\n", path, diff)
		}
		prompt += `Output:

		Generate a commit message that succinctly describes the modifications made within the identified code blocks.
		The commit message should accurately reflect the changes and provide enough context for others to understand the purpose of the commit.
		`

		commitMessage, err := generator.GenerateCommitMessage(prompt, model)
		if err != nil {
			log.Fatal(err)
		}

		if preview {

			var (
				reset   = "\033[0m"
				cYellow = "\033[33m"
			)
			fmt.Println(cYellow + commitMessage + reset)

			fmt.Println("Do you want to commit this message? (y/n)")
			var answer string
			fmt.Scanln(&answer)
			if answer != "y" && answer != "Y" {
				return
			}
		}

		if err := git.Commit(commitMessage); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("model", "m", "llama2", "model to use")
	rootCmd.PersistentFlags().BoolP("preview", "p", false, "preview the commit message")
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
