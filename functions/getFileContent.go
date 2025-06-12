package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileContent(workingDirectory, filePath string) (string, error) {
	fullPath := filepath.Join(workingDirectory, filePath)

	info, err := os.Stat(fullPath)
	if err != nil {
		return "", fmt.Errorf("cannot read '%s' as it is outside the permitted working directory", filePath)
	}

	if info.IsDir() {
		return "", fmt.Errorf("'%s' is not a file", filePath)
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
