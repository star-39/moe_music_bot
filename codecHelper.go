package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

//var embedFFMpeg = "OFF"
var ffmpegPath = "ffmpeg"

// opus only, it's the best, and Telegram supports it across all devices.
func transcodeAudio(fileName string, bitrate string) (string, error) {
	println(ffmpegPath)
	ffErr := checkFFMpeg()
	if ffErr != nil {
		return "", ffErr
	}

	tempDir := os.TempDir()
	fileBaseName := filepath.Base(fileName)
	outputFileName := filepath.Join(tempDir, "_temp_"+fileBaseName+".opus")

	cmd := exec.Command(ffmpegPath, "-i", fileName, "-c:a", "libopus",
		"-b:a", bitrate, "-y", outputFileName)
	err := cmd.Run()
	if err != nil {
		return "", err
	} else {
		return outputFileName, nil
	}
}

func checkFFMpeg() error {
	cmd := exec.Command(ffmpegPath, "-h")
	err := cmd.Run()
	if err != nil {
		return err
	} else {
		return nil
	}
}
