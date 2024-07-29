package utils

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"text/template"
)

var templates embed.FS

func IsFramework(framework string) bool {

	framework = strings.ToLower(framework)

	isFramework := []string{"gin", "fiber", "echo"}

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

func CreateFileFromTemplate(projectName, templateName, filePath, framework, repoUrl string) {
	fullPath := path.Join(projectName, filePath)
	fileContent, err := readTemplateFile(templateName)
	if err != nil {
		return
	}

	if repoUrl != "" {
		projectName = repoUrl
	}
	println(templateName)
	writeToFile(projectName, fullPath, fileContent, framework, repoUrl)
}

func readTemplateFile(templateName string) (string, error) {
	templatePath := path.Join("templates", templateName)
	content, err := templates.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file %s: %w", templatePath, err)
	}
	return string(content), nil
}

func writeToFile(projectName, filePath, content, framework, repoUrl string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %s", filePath, err)
	}

	defer file.Close()

	data := struct {
		ProjectName string
		Framework   string
		RepoUrl     string
	}{
		ProjectName: projectName,
		Framework:   framework,
		RepoUrl:     repoUrl,
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

func PromptForInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return ""
	}
	return strings.TrimSpace(input)
}

// PromptForInputWithValidation implement this later
func PromptForInputWithValidation(prompt string, validate func(string) bool) string {
	for {
		input := PromptForInput(prompt)
		if validate(input) {
			return input
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CleanRepoURL(url string) string {
	// Remove "https://", "http://", "www.", "https://www." if present
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "www.")
	url = strings.TrimPrefix(url, "https://www.")

	// Remove trailing "/"
	url = strings.TrimSuffix(url, "/")

	return url
}

// CopyTemplateFile This function is used to create stubborn templates eg. github actions
func CopyTemplateFile(libraryName, templateName, outputPath string) error {

	content, err := readTemplateFile(templateName)
	if err != nil {
		return fmt.Errorf("error reading template file: %w", err)
	}

	fullPath := path.Join(libraryName, outputPath)

	er := os.MkdirAll(path.Dir(fullPath), os.ModePerm)
	if er != nil {
		return fmt.Errorf("error creating directories: %w", er)
	}

	// Create the output file
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	// Write the content to the output file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to output file: %w", err)
	}

	// Verify file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Errorf("file was not created at %s", fullPath)
	}
	return nil
}

func CreateProjectDirectories(projectName string, dirs []string) error {
	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	return nil
}

func InstallDependencies(projectName, flag string, packages []string) {
	os.Chdir(projectName)

	log.Println("package installation started...")

	if len(packages) != 0 {

		for _, pkg := range packages {

			if err := RunCommand("go", "install", pkg); err != nil {
				log.Fatalf("Failed to install package %s: %s", pkg, err)
			}
		}
	}

	if flag == "p" {
		RunCommand("swag", "init")
	}

	RunCommand("go", "mod", "tidy")
	log.Println("All packages installed successfully")
}

func ValidateInputValue(inputName, inputValue string) string {
	for {
		if inputValue == "" {
			fmt.Printf("Error: %s cannot be empty\n", inputName)
		} else if strings.HasPrefix(inputValue, "-") {
			fmt.Printf("Error: %s cannot start with a hyphen (-)\n", inputName)
		} else {
			return inputValue
		}

		inputValue = PromptForInput(fmt.Sprintf("Please provide a valid %s: ", inputName))
	}
}
