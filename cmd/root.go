package cmd

import (
	"bufio"
	"fmt"
	"goeasycli/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	version     = "dev"
	projectName string
	framework   string
	repoUrl     string
	libraryName string
)

var rootCmd = &cobra.Command{
	Use:   "goeasycli",
	Short: "CLI tool to create a Gin project",
	Long: `A CLI tool to create web projects using different frameworks like:
	- Gin
	- Fiber
	- Echo`,
	Example: "goeasycli -p project_name -f  framework \neg. goeasycli -p fafa_shop_api -f gin ",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if projectName != "" && libraryName != "" {
			fmt.Println("Both project and library flags are present. Prioritizing project creation.")
			createProject()
		} else if projectName != "" {
			createProject()
		} else if libraryName != "" {
			createLibrary()
		} else {
			fmt.Println("Please specify either a project name (-p) or a library name (-l)")
			cmd.Usage()
			os.Exit(1)
		}
	},
}

/*
TODOS:
- add support for other orm  --orm -g/--gorm
-
*/

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Name of the project")
	rootCmd.PersistentFlags().StringVarP(&framework, "framework", "f", "", "web frameworks supported: (gin,fiber)")
	rootCmd.PersistentFlags().StringVarP(&libraryName, "library", "l", "", "Name of the library to create")
	rootCmd.PersistentFlags().StringVarP(&repoUrl, "repo", "r", "", "Repository URL for the library")
	// rootCmd.MarkPersistentFlagRequired("project")

}

func createProject() {
	dir, _ := os.Getwd()

	if strings.HasPrefix(projectName, "-") {
		fmt.Println("Error: Project name cannot start with a hyphen (-)")
		os.Exit(1)
	}

	if framework == "" {
		framework = utils.PromptForInput("Please choose a framework (gin,fibe,echo):")
	}

	exists := utils.IsFramework(framework)

	for {
		if !exists {
			fmt.Printf("framework entered '%s' does not exist.\n", framework)
			framework = utils.PromptForInput("Please choose a framework (gin,fibe,echo):")
			exists = utils.IsFramework(framework)
		} else {
			break
		}
	}

	fullpath := fmt.Sprintf("%s/%s", dir, projectName)

	// check if that directory already exist?
	for {
		if _, err := os.Stat(fullpath); err == nil {
			fmt.Printf("Project '%s' already exists. Please enter a new project name: ", projectName)
			reader := bufio.NewReader(os.Stdin)
			newName, _ := reader.ReadString('\n')
			projectName = strings.TrimSpace(newName)
			fullpath = fmt.Sprintf("%s/%s", dir, projectName)
		} else {
			break
		}
	}

	createProjectStructure()
	utils.OpenDirectory(fullpath)
}

func createLibrary() {
	dir, _ := os.Getwd()

	// Check if library name starts with a hyphen
	for strings.HasPrefix(libraryName, "-") {
		if libraryName != "" {
			fmt.Println("Error: Library name cannot start with a hyphen (-)")
		}
		libraryName = utils.PromptForInput("Please provide the library name: ")
	}

	if repoUrl == "" {
		repoUrl = utils.PromptForInput("Please provide the repository URL")
	}

	repoUrl = utils.CleanRepoURL(repoUrl)

	fullpath := filepath.Join(dir, libraryName)

	// Check if that directory already exists
	for {
		if _, err := os.Stat(fullpath); err == nil {
			fmt.Printf("Library '%s' already exists. Please enter a new library name: ", libraryName)

			libraryName = utils.PromptForInput("")

			// Check if the new library name starts with a hyphen
			if strings.HasPrefix(libraryName, "-") {
				fmt.Println("Error: Library name cannot start with a hyphen (-)")
				continue
			}

			fullpath = filepath.Join(dir, libraryName)
		} else {
			break
		}
	}

	createLibraryStructure(libraryName, repoUrl)
	utils.OpenDirectory(fullpath)

}

func createLibraryStructure(lName, repo string) {
	dirs := []string{
		"database",
		"email",
		".github/workflows",
		"loggers",
		"config",
	}

	for _, dir := range dirs {
		path := filepath.Join(libraryName, dir)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory %s: %s", path, err)
		}
	}

	createLibraryProject(libraryName, repoUrl)

}

func createLibraryProject(lName, repo string) {
	
	libraryFiles := map[string]string{
		"library/go.mod.tmpl": "go.mod",
		"library/emails.go.tmpl":  "email/emails.go",
		"shared/config.go.tmpl":   "config/config.go",
		"shared/database.go.tmpl": "database/database.go",
		"shared/zap.go.tmpl":      "loggers/zap.go",
	}

	for templateName, filePath := range libraryFiles {
		utils.CreateFileFromTemplate(lName, templateName, filePath, "", repo)
	}

	err := utils.CopyTemplateFile(lName, "library/workflow.tmpl", ".github/workflows/goeasycli_tag.yml")
	if err != nil {
		log.Fatalf("Failed to copy template file: %v", err)
	}

	os.Chdir(lName)
	utils.RunCommand("go", "mod", "tidy")
}

func createProjectStructure() {
	dirs := []string{
		"middlewares",
		"config",
		"routers",
		"models",
		"structs",
		"database",
		"controllers",
		"utils",
		"loggers",
		"templates",
		"seeds",
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory %s: %s", path, err)
		}
	}

	createProjectFiles(projectName, framework)

}

func createProjectFiles(projectName, framework string) {

	sharedFiles := map[string]string{
		"shared/env.tmpl":              ".env",
		"shared/middlewares.go.tmpl":   "middlewares/middlewares.go",
		"shared/config.go.tmpl":        "config/config.go",
		"shared/utils.go.tmpl":         "utils/utils.go",
		"shared/user_model.go.tmpl":    "models/user.go",
		"shared/go.mod.tmpl":           "go.mod",
		"shared/database.go.tmpl":      "database/database.go",
		"shared/template.go.tmpl":      "templates/templates.go",
		"shared/zap.go.tmpl":           "loggers/zap.go",
		"shared/sentry.go.tmpl":        "loggers/sentry.go",
		"shared/test_database.db.tmpl": "test_database.db",
		"shared/gitignore.tmpl":        ".gitignore",
		"shared/seeds.tmpl":            "seeds/seed.go",
	}

	frameworkFiles := map[string]map[string]string{
		"gin": {
			"gin/main.go.tmpl":              "main.go",
			"gin/health_controller.go.tmpl": "controllers/health.go",
			"gin/user_controller.go.tmpl":   "controllers/user.go",
			"gin/routers.go.tmpl":           "routers/routers.go",
			"gin/user_routers.go.tmpl":      "routers/user.go",
			"gin/health_router.go.tmpl":     "routers/health.go",
			"gin/responses.go.tmpl":         "utils/responses.go",
		},
		"fiber": {
			"fiber/user_controller.go.tmpl": "controllers/user.go",
			"fiber/routers.go.tmpl":         "routers/routers.go",
			"fiber/user_routers.go.tmpl":    "routers/user.go",
			"fiber/main.go.tmpl":            "main.go",
			"fiber/responses.go.tmpl":       "utils/responses.go",
		},
		"echo": {
			"echo/main.go.tmpl":              "main.go",
			"echo/health_controller.go.tmpl": "controllers/health.go",
			"echo/user_controller.go.tmpl":   "controllers/user.go",
			"echo/routers.go.tmpl":           "routers/routers.go",
			"echo/user_routers.go.tmpl":      "routers/user.go",
			"echo/health_router.go.tmpl":     "routers/health.go",
			"echo/responses.go.tmpl":         "utils/responses.go",
		},
	}

	if files, ok := frameworkFiles[framework]; ok {
		for templateName, filePath := range files {
			utils.CreateFileFromTemplate(projectName, templateName, filePath, framework, "")
		}
	} else {
		log.Fatalf("Unsupported framework: %s", framework)
	}

	for templateName, filePath := range sharedFiles {
		utils.CreateFileFromTemplate(projectName, templateName, filePath, framework, "")
	}

	installDependencies()
}

func installDependencies() {
	os.Chdir(projectName)
	utils.RunCommand("go", "mod", "tidy")

	// goGetPackages := []string{"github.com/gin-gonic/gin"}

	// for _, pkg := range goGetPackages {
	// 	if err := runCommand("go", "get", pkg); err != nil {
	// 		log.Fatalf("Failed to install package %s: %s", pkg, err)
	// 	}
	// }

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
