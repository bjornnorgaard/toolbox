package zeropad

import (
	"fmt"
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
		return fmt.Errorf("invalid root dir: %v", err)
	}

	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range files {
		if entry.IsDir() {
			err = CurrentDir(entry.Name())
			if err != nil {
				return fmt.Errorf("recursive call failed: %v", err)
			}
		}

		if !isImageFile(entry.Name()) {
			continue
		}

		newName := zeroPadFileName(entry.Name())
		newPath := filepath.Join(root, newName)

		err = os.Rename(filepath.Join(root, entry.Name()), newPath)
		if err != nil {
			return err
		}
	}

	return nil
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

func zeroPadFileName(fileName string) string {
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileNumber, err := strconv.Atoi(fileNameWithoutExt)
	if err != nil {
		err = fmt.Errorf("failed to convert string to int: %v", err)
		panic(err)
	}
	paddedName := fmt.Sprintf("%04d", fileNumber)
	return paddedName + filepath.Ext(fileName)
}
