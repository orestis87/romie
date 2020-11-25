package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// FolderExists reports whether the provided directory exists.
func FolderExists(path string) bool {
	if path == "" {
		log.Debug("Path is empty")
		return false
	}

	exists, err := AppFS.DirExists(path)
	if err != nil {
		log.Debug(fmt.Printf("Error: %s", err.Error()))
		return false
	}

	return exists
}

// FileExists reports whether the provided file exists.
func FileExists(path string) bool {
	if path == "" {
		log.Debug("Path is empty")
		return false
	}

	exists, err := AppFS.Exists(path)
	if err != nil {
		log.Debug(fmt.Printf("Error: %s", err.Error()))
		return false
	}

	if exists {
		isDir, _ := AppFS.IsDir(path)
		return !isDir
	}

	return false
}
