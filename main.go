package main

import (
	"bufio"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const megaBytes int64 = 1000000

var transcodeToOpus = false
var transcodeBitrate = "192k"

func main() {
	apiToken := ""
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		println("Please enter your Telegram BOT API token:")
		apiToken, _ = reader.ReadString('\n')
		apiToken = strings.TrimSuffix(apiToken, "\n")
	} else {
		apiToken = os.Args[1]
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  apiToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	println("Server OK. You are authorized as: ", b.Me.Username)
	println("The bot is now started.")

	b.Handle("/set", func(m *tb.Message) {
		commandSet(m, b)
	})
	b.Handle("/up", func(m *tb.Message) {
		commandUp(m, b)
	})
	b.Handle("/help", func(m *tb.Message) {
		commandHelp(m, b)
	})

	b.Start()
}

func commandUp(m *tb.Message, b *tb.Bot) {
	dir := m.Payload
	println("directory is : ", dir)
	audioFiles, _ := GetAudioFiles(dir)
	for _, audioFile := range audioFiles {
		println("hit : ", audioFile)
		audioTag, err := ReadFileAndTag(audioFile)
		if err != nil {
			b.Send(m.Chat, "error on file : "+audioFile)
			continue
		}
		fileStat, _ := os.Stat(audioFile)
		//Transcode
		if fileStat.Size() > 50*megaBytes || transcodeToOpus {
			var transErr error = nil
			audioFile, transErr = transcodeAudio(audioFile, transcodeBitrate)
			if transErr != nil {
				b.Send(m.Chat, "ffmpeg error on file : "+audioFile)
				b.Send(m.Chat, transErr.Error())
				continue
			}
			defer os.Remove(audioFile)
		}

		audioSendable := &tb.Audio{
			File:      tb.FromDisk(audioFile),
			Title:     audioTag.title,
			Performer: audioTag.performer,
			Caption: "#" + strings.Join(strings.Fields(audioTag.album), "_") + "\n" +
				"#" + strings.Join(strings.Fields(audioTag.performer), ""),
		}

		_, err2 := b.Send(m.Chat, audioSendable)
		if err2 != nil {
			b.Send(m.Chat, "error sending : "+audioFile)
		}
	}
}

func commandSet(m *tb.Message, b *tb.Bot) {
	args := strings.Fields(m.Payload)
	if len(args) == 0 {
		b.Send(m.Chat, strconv.FormatBool(transcodeToOpus)+transcodeBitrate)
		return
	}
	if args[0] == "transcode" {
		if args[1] == "yes" {
			transcodeToOpus = true
		} else {
			transcodeToOpus = false
		}
	} else if args[0] == "bitrate" {
		transcodeBitrate = args[1]
	}
}

func commandHelp(m *tb.Message, b *tb.Bot) {
	b.Send(m.Chat,
		"/up <directory>\n"+
			"Upload all musics from directory (recursively)\n\n"+
			"/set\n"+
			"Print current settings\n\n"+
			"/set transcode [yes|no]\n"+
			"Transcode to opus 192k\n\n"+
			"/set bitrate <bitrate>\n"+
			"bitrate should end with 'k', eg. 256k")
}
