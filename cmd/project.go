package cmd

import (
	"fmt"
	"goeasycli/utils"
	"log"
	"os"
	"path/filepath"
)

func createProject(projectName, framework string) {
	dir, _ := os.Getwd()

	projectName = utils.ValidateInputValue("Project", projectName)

	fullpath := filepath.Join(dir, projectName)

	// check if that directory already exist?
	if _, err := os.Stat(fullpath); err == nil {
		fmt.Printf("Project '%s' already exists.", projectName)
		projectName = utils.PromptForInput("Please enter a new project name: ")
		projectName = utils.ValidateInputValue("Project", projectName)
		fullpath = filepath.Join(dir, projectName)
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

	createProjectFiles(projectName, framework)
	utils.OpenDirectory(fullpath)
}

func createProjectFiles(projectName, framework string) {

	dirs := []string{
		"middlewares",
		"config",
		"routers",
		"models",
		"structs",
		"database",
		"controllers",
		"utils",
		"logger",
		"templates",
		"seeds",
	}

	err := utils.CreateProjectDirectories(projectName, dirs)

	if err != nil {
		log.Printf("Error creating project directories: %v", err)
	}

	sharedFiles := map[string]string{
		"shared/env.tmpl":                ".env",
		"shared/middlewares.go.tmpl":     "middlewares/middlewares.go",
		"shared/config.go.tmpl":          "config/config.go",
		"shared/utils.go.tmpl":           "utils/utils.go",
		"shared/user_model.go.tmpl":      "models/user.go",
		"shared/go.mod.tmpl":             "go.mod",
		"shared/database.go.tmpl":        "database/database.go",
		"shared/template.go.tmpl":        "templates/templates.go",
		"shared/zap.go.tmpl":             "logger/zap.go",
		"shared/sentry.go.tmpl":          "logger/sentry.go",
		"shared/test_database.db.tmpl":   "test_database.db",
		"shared/gitignore.tmpl":          ".gitignore",
		"shared/seeds.tmpl":              "seeds/seed.go",
		"shared/structs.go.tmpl":         "structs/structs.go",
		"shared/dockerignore.tmpl":       ".dockerignore",
		"shared/Dockerfile.tmpl":         "Dockerfile",
		"shared/docker-compose.yml.tmpl": "docker-compose.yml",
		"shared/Makefile.tmpl":           "Makefile",
		"shared/Readme.md.tmpl":          "Readme.md",
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
			"fiber/main.go.tmpl":              "main.go",
			"fiber/health_controller.go.tmpl": "controllers/health.go",
			"fiber/user_controller.go.tmpl":   "controllers/user.go",
			"fiber/routers.go.tmpl":           "routers/routers.go",
			"fiber/user_routers.go.tmpl":      "routers/user.go",
			"fiber/health_router.go.tmpl":     "routers/health.go",
			"fiber/responses.go.tmpl":         "utils/responses.go",
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

	utils.InstallDependencies(projectName)
}
