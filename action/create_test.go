package action

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/urfave/cli/v3"
)

// MockCmd is a mock for exec.Cmd
type MockCmd struct {
	shouldFail bool
}

func (m *MockCmd) Run() error {
	if m.shouldFail {
		return fmt.Errorf("mock command failed")
	}
	return nil
}

// MockExecCommand creates a mock exec.Command function
func MockExecCommand(command string, args ...string) *exec.Cmd {
	cmd := &exec.Cmd{
		Path: command,
		Args: append([]string{command}, args...),
	}
	return cmd
}
func buildMinimalTestCommand() *cli.Command {
	// reset the help flag because tests may have set it
	cli.HelpFlag.(*cli.BoolFlag).Name = "test"
	return &cli.Command{Writer: io.Discard}
}

// TestCreateProject_EmptyProjectName tests the case when project name is empty
func TestCreateProject_EmptyProjectName(t *testing.T) {
	projectName := ""
	err := createProject(projectName)

	// Verify that the function returns nil and prints the correct message
	if err.Error() != "please specifiy the new ProjectName" {
		t.Error("Expected error message to contain 'please specifiy the new ProjectName'")
	}
}

func TestCreateProject_ValidProjectName(t *testing.T) {
	// Create a temp dir as workdir for testing.
	workDir, err := os.MkdirTemp("", "test-")
	defer os.RemoveAll(workDir)
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	os.Chdir(workDir)
	println("workDir: %s", workDir)
	projectName := "test-project"
	err = createProject(projectName)

	// Verify that the function returns nil
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
