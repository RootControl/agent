package functions

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetFilesInfo(t *testing.T) {
	// Create temporary working directory
	workingDir, err := os.MkdirTemp("", "test-working-dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(workingDir)

	// Create a subdirectory to test
	subDir := "testdata"
	fullSubDir := filepath.Join(workingDir, subDir)
	err = os.Mkdir(fullSubDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Create a file inside the subdir
	file1 := filepath.Join(fullSubDir, "README.md")
	err = os.WriteFile(file1, []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create another directory inside the subdir
	nestedDir := filepath.Join(fullSubDir, "src")
	err = os.Mkdir(nestedDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Create another file
	file2 := filepath.Join(fullSubDir, "package.json")
	err = os.WriteFile(file2, []byte(`{}`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Run the function
	output, err := GetFilesInfo(workingDir, subDir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check output
	expected := []string{
		"- README.md: file_size=5 bytes, is_dir=false",
		"- src: file_size=128 bytes, is_dir=true", // exact dir size may vary by system
		"- package.json: file_size=2 bytes, is_dir=false",
	}

	for _, exp := range expected {
		if !strings.Contains(output, exp[:strings.Index(exp, ":")]) {
			t.Errorf("expected entry not found: %s", exp)
		}
	}
}
