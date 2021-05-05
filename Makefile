GO=go


all: init build

init:
	$(GO) mod init moe_music_bot
	$(GO) get -u gopkg.in/tucnak/telebot.v2
	$(GO) get -u https://github.com/dhowden/tag

build:
	$(GO) build

with-ffmpeg:
	$(GO) build -tags embed_ffmpeg

clean:
	rm moe_music_bot
