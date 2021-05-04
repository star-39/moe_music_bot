package main

import (
	"os/exec"
	"path/filepath"
)

// opus only, it's the best, and Telegram supports it across all devices.
func transcodeAudio(fileName string, bitrate string) (string, error) {
	fileDir := filepath.Dir(fileName)
	fileBaseName := filepath.Base(fileName)
	outputFileName := filepath.Join(fileDir, "_temp_"+fileBaseName+".opus")

	cmd := exec.Command("ffmpeg", "-i", fileName, "-c:a", "libopus",
		"-b:a", bitrate, "-y", outputFileName)
	err := cmd.Run()
	if err != nil {
		return "", err
	} else {
		return outputFileName, nil
	}
}
