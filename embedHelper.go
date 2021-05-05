// +build embed_ffmpeg

package main

import (
	_ "embed"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
)

//go:embed ffmpeg
var ffmpeg_bin []byte

var signalChan = make(chan os.Signal)

func init() {
	print("This binary is built with embedded FFMpeg, extracting to ")
	println(path.Join(os.TempDir(), "_temp_ffmpeg"))
	ioutil.WriteFile(path.Join(os.TempDir(), "_temp_ffmpeg"), ffmpeg_bin, 0755)

	signal.Notify(signalChan, os.Interrupt)
	go onExit()
}

func onExit() {
	select {
	case <-signalChan:
		println("Exiting... Doing clean up...")
		os.Remove(path.Join(os.TempDir(), "_temp_ffmpeg"))
		os.Exit(0)
	}
}
