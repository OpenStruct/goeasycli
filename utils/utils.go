package utils

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"text/template"
)

var templates embed.FS

func IsFramework(framework string) bool {

	framework = strings.ToLower(framework)

	isFramework := []string{"gin", "fiber"}

	exists := slices.Contains(isFramework, framework)

	return exists
}

func OpenDirectory(projectPath string) {
	var cmd *exec.Cmd

	// Define a function to check if a command exists
	checkCommandExists := func(command string) bool {
		_, err := exec.LookPath(command)
		return err == nil
	}

	// List of applications to check
	apps := []string{"code", "cursor", "atom", "goland"}

	// Find the first available application
	var availableApp string
	for _, app := range apps {
		if checkCommandExists(app) {
			availableApp = app
			break
		}
	}

	switch runtime.GOOS {
	case "windows":
		if availableApp != "" {
			cmd = exec.Command("cmd", "/c", availableApp, projectPath)
		} else {
			cmd = exec.Command("explorer", projectPath)
		}
	case "darwin":
		if availableApp != "" {
			switch availableApp {
			case "code":
				cmd = exec.Command("open", "-a", "Visual Studio Code", projectPath)
			case "cursor":
				cmd = exec.Command("open", "-a", "Cursor", projectPath)
			case "atom":
				cmd = exec.Command("open", "-a", "Atom", projectPath)
			case "goland":
				cmd = exec.Command("open", "-a", "GoLand", projectPath)
			}
		} else {
			cmd = exec.Command("open", projectPath)
		}
	case "linux":
		if availableApp != "" {
			cmd = exec.Command(availableApp, projectPath)
		} else {
			cmd = exec.Command("xdg-open", projectPath)
		}
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
}

func CreateFileFromTemplate(projectName, templateName, filePath string) {
	fullPath := filepath.Join(projectName, filePath)
	fileContent, err := readTemplateFile(templateName)
	if err != nil {
		return
	}
	writeToFile(projectName, fullPath, fileContent)
}

func readTemplateFile(templateName string) (string, error) {
	templatePath := filepath.Join("templates", templateName)
	content, err := templates.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}
	return string(content), nil
}

func writeToFile(projectName, filePath, content string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %s", filePath, err)
	}

	defer file.Close()

	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	// Parse the content as a template
	tmpl, err := template.New("file").Parse(string(content))
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
	}

	// Execute the template and write to the file
	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
}

func SetTemplatesFS(fs embed.FS) {
	templates = fs
}

func PromptForFramework() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please choose a framework (gin,fiber):")
	framework, _ := reader.ReadString('\n')
	framework = strings.TrimSpace(framework)
	return framework
}

func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
