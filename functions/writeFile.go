package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(workingDirectory, filePath, content string) error {
	fullPath := filepath.Join(workingDirectory, filePath)

	_, err := os.Stat(workingDirectory)
	if err != nil {
		return fmt.Errorf("cannot read '%s' as it is outside the permitted working directory", filePath)
	}

	return os.WriteFile(fullPath, []byte(content), 0644)
}
