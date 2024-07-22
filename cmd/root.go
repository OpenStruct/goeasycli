package cmd

import (
	"fmt"
	"goeasycli/utils"
	"os"
	"runtime"

	"github.com/spf13/cobra"
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
		uninstallFlag, _ := cmd.Flags().GetBool("uninstall")
		if uninstallFlag {
			uninstall()
			return
		}

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
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.Flags().BoolP("uninstall", "u", false, "Uninstall goeasycli")
	rootCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Name of the project")
	rootCmd.PersistentFlags().StringVarP(&framework, "framework", "f", "", "Web frameworks supported: (gin,fiber,echo)")
	rootCmd.PersistentFlags().StringVarP(&libraryName, "library", "l", "", "Name of the library to create")
	rootCmd.PersistentFlags().StringVarP(&repoUrl, "repo", "r", "", "Repository url for the library")
}

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"u"},
	Short:   "Uninstall goeasycli",
	Long:    `Uninstall goeasycli from your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall()
	},
}

func uninstall() {
	var binaryPath string

	switch runtime.GOOS {
	case "windows":
		binaryPath = os.Getenv("PROGRAMFILES") + "\\goeasycli\\goeasycli.exe"
	case "darwin", "linux":
		binaryPath = "/usr/local/bin/goeasycli"
	default:
		fmt.Printf("Unsupported operating system: %s\n", runtime.GOOS)
		os.Exit(1)
	}

	// Check if the binary exists
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		fmt.Println("goeasycli is not installed.")
		return
	}

	fmt.Println("Uninstalling goeasycli...")

	var err error
	switch runtime.GOOS {
	case "windows":
		err = utils.RunCommand("cmd", "/C", "del", binaryPath)
	case "darwin", "linux":
		err = utils.RunCommand("sudo", "rm", binaryPath)
	}

	if err != nil {
		fmt.Printf("Failed to uninstall goeasycli: %v\n", err)
		os.Exit(1)
	}

	// Remove the directory on windows
	if runtime.GOOS == "windows" {
		dirPath := os.Getenv("PROGRAMFILES") + "\\goeasycli"
		err = os.RemoveAll(dirPath)
		if err != nil {
			fmt.Printf("Failed to remove directory: %v\n", err)
		}
	}

	fmt.Println("goeasycli has been successfully uninstalled.")

	// Remove from PATH for Windows
	if runtime.GOOS == "windows" {
		fmt.Println("Please note: You may need to manually remove goeasycli from your system PATH.")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
