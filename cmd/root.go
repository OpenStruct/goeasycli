package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	version     = "dev"
	projectName string
	framework   string
	libraryName string
	repoUrl     string
)

var longDescription = `A CLI tool to create web projects using different frameworks like:
	- Gin
	- Fiber
	- Echo
	
It can also be used to create libraries.

Usage:
  Create a new project:
    goeasycli -p project_name -f framework
    e.g: goeasycli  -p fafa_shop_api -f gin

  Create a new library:
    goeasycli -l library_name -r repo_url
    e.g: goeasycli -l my_awesome_fafa_lib -r https://github.com/heygoeasycli/my_awesome_fafa_lib`

var rootCmd = &cobra.Command{
	Use:     "goeasycli",
	Short:   "CLI tool to create Go projects and libraries",
	Long:    longDescription,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if projectName != "" && libraryName != "" {
			fmt.Println("Both project and library flags are present. Prioritizing project creation.")
			createProject(projectName, framework)
		} else if projectName != "" {
			createProject(projectName, framework)
		} else if libraryName != "" {
			createLibrary(libraryName, repoUrl)
		} else {
			fmt.Println("Please specify either a project name (-p) or a library name (-l)")
			cmd.Usage()
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Name of the project")
	rootCmd.PersistentFlags().StringVarP(&framework, "framework", "f", "", "Web frameworks supported: (gin,fiber,echo)")
	rootCmd.PersistentFlags().StringVarP(&libraryName, "library", "l", "", "Name of the library to create")
	rootCmd.PersistentFlags().StringVarP(&repoUrl, "repo", "r", "", "Repository url for the library")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
