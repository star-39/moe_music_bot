// +build embed_ffmpeg

package main

import (
	_ "embed"
	"os"
	"os/signal"
	"path"
	"runtime"
)

//go:embed ffmpeg
var ffmpeg_bin []byte

var signalChan = make(chan os.Signal)

func init() {
	if runtime.GOOS == "windows" {
		ffmpegPath = path.Join(os.TempDir(), "_temp_ffmpeg.exe")
	} else {
		ffmpegPath = path.Join(os.TempDir(), "_temp_ffmpeg")
	}

	println("This binary is built with embedded FFMpeg, extracting to ", ffmpegPath)

	os.WriteFile(ffmpegPath, ffmpeg_bin, 0755)

	signal.Notify(signalChan, os.Interrupt)
	go onExit()
}

func onExit() {
	select {
	case <-signalChan:
		println("Exiting... Doing clean up...")
		os.Remove(ffmpegPath)
		os.Exit(0)
	}
}
