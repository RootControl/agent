package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFilesInfo(workingDirectory, directory string) (string, error) {
	fullPath := filepath.Join(workingDirectory, directory)

	info, err := os.Stat(fullPath)
	if err != nil {
		return "", fmt.Errorf("cannot list '%s' as it is outside the permitted working directory", directory)
	}

	if !info.IsDir() {
		return "", fmt.Errorf("'%s' is not a directory", directory)
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return "", err
	}

	var result string

	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			continue
		}

		result += fmt.Sprintf("- %s: file_size=%d bytes, is_dir=%v\n",
			entry.Name(),
			entryInfo.Size(),
			entry.IsDir())
	}

	return result, nil
}
