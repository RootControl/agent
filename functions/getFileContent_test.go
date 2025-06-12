package functions

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	// Create temporary working directory
	workingDir, err := os.MkdirTemp("", "test-working-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workingDir)

	// Create a file to test
	filePath := "test.md"
	fullPath := filepath.Join(workingDir, filePath)
	err = os.WriteFile(fullPath, []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Call the function
	content, err := GetFileContent(workingDir, filePath)
	if err != nil {
		t.Fatal(err)
	}

	// Check the result
	expectedContent := "hello"
	if content != expectedContent {
		t.Errorf("Expected content '%s', got '%s'", expectedContent, content)
	}

	// Check the file permissions
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		t.Fatal(err)
	}
	expectedPermissions := os.FileMode(0644)
	if fileInfo.Mode() != expectedPermissions {
		t.Errorf("Expected file permissions %s, got %s", expectedPermissions, fileInfo.Mode())
	}
}
