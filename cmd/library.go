package cmd

import (
	"fmt"
	"goeasycli/utils"
	"log"
	"os"
	"path/filepath"
)

func createLibrary(libraryName, repo string) {
	dir, _ := os.Getwd()

	// Check if library name starts with a hyphen
	libraryName = utils.ValidateInputValue("library", libraryName)

	fullpath := filepath.Join(dir, libraryName)

	// Check if that directory already exists

	if _, err := os.Stat(fullpath); err == nil {
		fmt.Printf("Library '%s' already exists.", libraryName)
		libraryName = utils.PromptForInput("Please enter a new library name: ")
		libraryName = utils.ValidateInputValue("library", libraryName)
		fullpath = filepath.Join(dir, libraryName)
	}

	if repo == "" {
		repo = utils.PromptForInput("Please provide the repository URL")
	}

	repo = utils.CleanRepoURL(repo)

	createLibraryProject(libraryName, repo)
	utils.OpenDirectory(fullpath)

}

func createLibraryProject(lName, repo string) {
	dirs := []string{
		"database",
		"email",
		".github/workflows",
		"loggers",
		"config",
	}

	err := utils.CreateProjectDirectories(lName, dirs)
	if err != nil {
		log.Printf("Error creating project directories: %v", err)
	}

	libraryFiles := map[string]string{
		"library/go.mod.tmpl":     "go.mod",
		"library/emails.go.tmpl":  "email/emails.go",
		"shared/config.go.tmpl":   "config/config.go",
		"shared/database.go.tmpl": "database/database.go",
		"shared/zap.go.tmpl":      "loggers/zap.go",
	}

	for templateName, filePath := range libraryFiles {
		utils.CreateFileFromTemplate(lName, templateName, filePath, "", repo)
	}

	err = utils.CopyTemplateFile(lName, "library/workflow.tmpl", ".github/workflows/goeasycli_tag.yml")
	if err != nil {
		log.Fatalf("Failed to copy template file: %v", err)
	}

	utils.InstallDependencies(lName)
}
