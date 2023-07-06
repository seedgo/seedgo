package action

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const originProjectname string = "seedgo-skeleton"

var ignoreDireList = []string{".git"}

func sliceContain(l []string, k string) bool {
	for _, i := range l {
		if i == k {
			return true
		}
	}

	return false
}

func CreateProject(projectName string) error {
	fmt.Println("creating project: " + projectName)
	// 1. clone the repository
	repo := "https://github.com/seedgo/seedgo-skeleton"
	cmd := exec.Command("git", "clone", repo, projectName)
	fmt.Println(cmd.String())
	err := cmd.Run()

	// 2. replace the projectName
	err = filepath.WalkDir(projectName, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if sliceContain(ignoreDireList, info.Name()) {
				return filepath.SkipDir
			}
		} else {
			f, err := os.Open(path)
			if err != nil {
				fmt.Printf("open File %s err: %s\n:", path, err.Error())
			}
			defer f.Close()
			buf, err := io.ReadAll(f)
			if err != nil {
				fmt.Printf("read File %s err: %s\n:", path, err.Error())
			}

			content := string(buf)
			// Check if the file contains the string "".
			stringFound := strings.Contains(content, originProjectname)

			// If the string is found, replace it with "jaylin".
			if stringFound {
				content = strings.Replace(content, originProjectname, projectName, -1)
			}
			// Empty the original file.
			err = os.Truncate(path, 0)
			if err != nil {
				fmt.Printf("truncate file %s err: %s\n", path, err.Error())
			}

			err = os.WriteFile(path, []byte(content), 0644)
			if err != nil {
				fmt.Printf("write file %s err: %s\n", path, err.Error())
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("finished create project: " + projectName)
	return err
}
