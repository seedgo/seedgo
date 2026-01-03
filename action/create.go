package action

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

const originProjectname string = "seedgo-skeleton"

func CreateProject(ctx context.Context, command *cli.Command) error {
	projectName := command.Args().Get(0)
	return createProject(projectName)
}

func createProject(projectName string) error {
	if len(projectName) == 0 {
		return fmt.Errorf("please specifiy the new ProjectName")
	}

	// Validate project name to prevent command injection
	if !isValidProjectName(projectName) {
		return fmt.Errorf("invalid project name: %s", projectName)
	}

	fmt.Println("creating project: " + projectName)
	// 1. clone the repository
	repo := "https://github.com/seedgo/seedgo-skeleton"
	cmd := exec.Command("git", "clone", repo, projectName)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Download project template failed %s", err.Error())
	}

	// 2. delete the .git dir if exist
	err = os.RemoveAll(filepath.Join(projectName, ".git"))
	if err != nil {
		return fmt.Errorf("delete .git dir failed: %s", err.Error())
	}

	// 3. replace the projectName
	err = filepath.WalkDir(projectName, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open File %s err: %s\n:", path, err.Error())
		}

		defer f.Close()
		buf, err := io.ReadAll(f)
		if err != nil {
			return fmt.Errorf("read File %s err: %s\n:", path, err.Error())
		}

		content := string(buf)
		// Check if the file contains the string originProjectName.
		stringFound := strings.Contains(content, originProjectname)

		// If the string is found, replace it with projectName.
		if stringFound {
			content = strings.Replace(content, originProjectname, projectName, -1)
		}
		// Empty the original file.
		err = os.Truncate(path, 0)
		if err != nil {
			fmt.Printf("truncate file %s err: %s\n", path, err.Error())
		}

		// Write the content to the file
		err = os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("write file %s err: %s\n", path, err.Error())
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("finished create project: " + projectName)
	return err
}

func isValidProjectName(name string) bool {
	// Only allow alphanumeric characters, hyphens, underscores, and dots
	for _, r := range name {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.') {
			return false
		}
	}
	return len(name) > 0 && !strings.Contains(name, "..") && !strings.HasPrefix(name, "/") && !strings.Contains(name, "/../")
}
