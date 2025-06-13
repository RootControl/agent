package functions

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	workingDir, err := os.MkdirTemp("", "test-working-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workingDir)

	err = WriteFile(workingDir, "test.md", "hello")
	if err != nil {
		t.Fatal(err)
	}

	content, err := GetFileContent(workingDir, "test.md")
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "hello" {
		t.Errorf("Expected file content 'hello', got '%s'", string(content))
	}
}
