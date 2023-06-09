package zeropad

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CurrentDir(root string) error {
	if len(root) == 0 {
		return fmt.Errorf("empty root dir")
	}

	_, err := os.Stat(root)
	if err != nil {
		return err
	}

	var (
		paddedPath string
		newPath    string
	)

	return filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !isImageFile(path) {
			return err
		}

		paddedPath, err = zeroPadFileName(info.Name())
		if err != nil {
			return err
		}

		newPath = strings.ReplaceAll(path, info.Name(), paddedPath)
		return os.Rename(path, newPath)
	})
}

func isImageFile(fileName string) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif"} // Add more image extensions if needed
	lowerFileName := strings.ToLower(fileName)

	for _, ext := range extensions {
		if strings.HasSuffix(lowerFileName, ext) {
			return true
		}
	}

	return false
}

func zeroPadFileName(fileName string) (string, error) {
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileNumber, err := strconv.Atoi(fileNameWithoutExt)
	if err != nil {
		return "", fmt.Errorf("failed to convert string to int: %v", err)
	}
	paddedName := fmt.Sprintf("%04d", fileNumber)
	return paddedName + filepath.Ext(fileName), nil
}
