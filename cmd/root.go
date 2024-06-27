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

var projectName string
var framework string
var rootCmd = &cobra.Command{
	Use:   "gin-cli",
	Short: "CLI tool to create a Gin project",
	Long: `A CLI tool to create web projects using different frameworks like:
	- Gin
	- Fiber`,
	Example: "gostartpro -p project_name -f  framework \neg. gostartpro -p love_match_api -f gin ",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		if framework == "" {
			framework = utils.PromptForFramework()
		}

		exists := utils.IsFramework(framework)

		for {
			if !exists {
				fmt.Printf("framework entered '%s' does not exist.\n", framework)
				framework = utils.PromptForFramework()
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
	rootCmd.MarkPersistentFlagRequired("project")

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
		"shared/.env.tmpl":             ".env",
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
	}

	if files, ok := frameworkFiles[framework]; ok {
		for templateName, filePath := range files {
			utils.CreateFileFromTemplate(projectName, templateName, filePath)
		}
	} else {
		log.Fatalf("Unsupported framework: %s", framework)
	}

	for templateName, filePath := range sharedFiles {
		utils.CreateFileFromTemplate(projectName, templateName, filePath)
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
