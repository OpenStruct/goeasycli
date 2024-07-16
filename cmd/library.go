package cmd

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func addLibraryToProject() error {
// 	// Check if we're in a Go project
// 	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
// 		return fmt.Errorf("go.mod file not found. Are you in a Go project directory?")
// 	}

// 	// Add the library using go get
// 	cmd := exec.Command("go", "get", repoURL)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		return fmt.Errorf("failed to add library: %v", err)
// 	}

// 	fmt.Printf("Successfully added library %s from %s\n", libraryName, repoURL)
// 	return nil
// }