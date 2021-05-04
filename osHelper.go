package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetAudioFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".m4a") ||
				strings.HasSuffix(path, ".flac") ||
				strings.HasSuffix(path, ".opus") ||
				strings.HasSuffix(path, ".mp3") {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		return nil, errors.New("Error listing directory!")
	} else {
		return files, nil
	}
}
